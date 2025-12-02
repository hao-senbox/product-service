package product

import (
	"context"
	"errors"
	"fmt"
	"log"
	"product-service/internal/shared/ports"
	"product-service/internal/topic"
	"product-service/pkg/uploader"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req *CreateProductRequest) (string, error)
	GetAllProducts(ctx context.Context) ([]*ProductResponse, error)
	GetProduct(ctx context.Context, id string) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, req *UpdateProductRequest, id string) error
	DeleteProduct(ctx context.Context, id string) error
}

type productService struct {
	productRepostitory ProductRepository
	folderRepository   ports.FolderRepository
	topicService       topic.TopicService
	imageService       uploader.ImageService
}

func NewProductService(productRepostitory ProductRepository, folderRepository ports.FolderRepository, topicService topic.TopicService, imageService uploader.ImageService) ProductService {
	return &productService{
		productRepostitory: productRepostitory,
		folderRepository:   folderRepository,
		topicService:       topicService,
		imageService:       imageService,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *CreateProductRequest) (string, error) {

	if req.ProductName == "" {
		return "", errors.New("product name is required")
	}

	if req.OriginPriceStore == 0 {
		return "", errors.New("origin price store is required")
	}

	if req.OriginPriceService == 0 {
		return "", errors.New("origin price service is required")
	}

	if req.FolderID == "" {
		return "", errors.New("category id is required")
	}

	if req.TopicID == "" {
		return "", errors.New("topic id is required")
	}

	if req.CoverImage == "" {
		return "", errors.New("cover image is required")
	}

	folderObjectID, err := primitive.ObjectIDFromHex(req.FolderID)
	if err != nil {
		return "", err
	}

	topicObjectID, err := primitive.ObjectIDFromHex(req.TopicID)
	if err != nil {
		return "", err
	}

	ID := primitive.NewObjectID()

	QRCocde := fmt.Sprintf("SENBOX.ORG[PRODUCT]:%s", ID.Hex())

	product := &Product{
		ID:                 ID,
		ProductName:        req.ProductName,
		OriginPriceStore:   req.OriginPriceStore,
		OriginPriceService: req.OriginPriceService,
		ProductDescription: req.ProductDescription,
		CoverImage:         req.CoverImage,
		TopicID:            topicObjectID,
		FolderID:           folderObjectID,
		QRCode:             QRCocde,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	id, err := s.productRepostitory.CreateProduct(ctx, product)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *productService) GetAllProducts(ctx context.Context) ([]*ProductResponse, error) {

	res, err := s.productRepostitory.GetAllProducts(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return []*ProductResponse{}, nil
		}
		return nil, err
	}

	var products []*ProductResponse

	for _, product := range res {

		folder, err := s.folderRepository.GetFolder(ctx, product.FolderID)
		if err != nil {
			log.Println("Error getting folder:", err)
		}

		topic, err := s.topicService.GetTopicByID(ctx, product.TopicID.Hex())
		if err != nil {
			log.Println("Error getting topic:", err)
		}

		topicResp := &Topic{}
		if topic != nil {
			topicResp.ID = topic.ID
			topicResp.Name = topic.Name
		}

		folderResp := Folder{}
		if folder != nil {
			folderResp = Folder{
				ID:   folder.ID.Hex(),
				Name: folder.Name,
			}
		}

		var image string
		if product.CoverImage != "" {
			img, err := s.imageService.GetImageKey(ctx, product.CoverImage)
			if err != nil {
				log.Println("Error getting image key:", err)
			} else if img != nil {
				image = img.Url
			}
		}

		products = append(products, &ProductResponse{
			ID:                 product.ID,
			ProductName:        product.ProductName,
			OriginPriceStore:   product.OriginPriceStore,
			OriginPriceService: product.OriginPriceService,
			ProductDescription: product.ProductDescription,
			CoverImage:         image,
			Topic:              topicResp,
			Folder:             folderResp,
			QRCode:             product.QRCode,
			CreatedAt:          product.CreatedAt,
			UpdatedAt:          product.UpdatedAt,
		})

	}

	return products, nil
}

func (s *productService) GetProduct(ctx context.Context, id string) (*ProductResponse, error) {

	idObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	product, err := s.productRepostitory.GetProduct(ctx, idObjectID)
	if err != nil {
		return nil, err
	}

	folder, err := s.folderRepository.GetFolder(ctx, product.FolderID)
	if err != nil {
		log.Println("Error getting folder:", err)
	}

	topic, err := s.topicService.GetTopicByID(ctx, product.TopicID.Hex())
	if err != nil {
		log.Println("Error getting topic:", err)
	}

	topicResp := &Topic{}
	if topic != nil {
		topicResp.ID = topic.ID
		topicResp.Name = topic.Name
	}

	folderResp := Folder{}
	if folder != nil {
		folderResp = Folder{
			ID:   folder.ID.Hex(),
			Name: folder.Name,
		}
	}

	var image string
	if product.CoverImage != "" {
		img, err := s.imageService.GetImageKey(ctx, product.CoverImage)
		if err != nil {
			log.Println("Error getting image key:", err)
		} else if img != nil {
			image = img.Url
		}
	}

	return &ProductResponse{
		ID:                 product.ID,
		ProductName:        product.ProductName,
		OriginPriceStore:   product.OriginPriceStore,
		OriginPriceService: product.OriginPriceService,
		ProductDescription: product.ProductDescription,
		CoverImage:         image,
		Topic:              topicResp,
		Folder:             folderResp,
		QRCode:             product.QRCode,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          product.UpdatedAt,
	}, nil

}

func (s *productService) UpdateProduct(ctx context.Context, req *UpdateProductRequest, id string) error {

	idObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	product, err := s.productRepostitory.GetProduct(ctx, idObjectID)
	if err != nil {
		return err
	}

	if req.ProductName != "" {
		product.ProductName = req.ProductName
	}

	if req.OriginPriceStore != 0 {
		product.OriginPriceStore = req.OriginPriceStore
	}

	if req.OriginPriceService != 0 {
		product.OriginPriceService = req.OriginPriceService
	}

	if req.ProductDescription != "" {
		product.ProductDescription = req.ProductDescription
	}

	if req.CoverImage != "" {
		product.CoverImage = req.CoverImage
	}

	if req.TopicID != "" {
		topicObjectID, err := primitive.ObjectIDFromHex(req.TopicID)
		if err != nil {
			return err
		}
		product.TopicID = topicObjectID
	}

	if req.FolderID != "" {
		folderObjectID, err := primitive.ObjectIDFromHex(req.FolderID)
		if err != nil {
			return err
		}
		product.FolderID = folderObjectID
	}

	productData := &Product{
		ID:                 product.ID,
		ProductName:        product.ProductName,
		OriginPriceStore:   product.OriginPriceStore,
		OriginPriceService: product.OriginPriceService,
		ProductDescription: product.ProductDescription,
		CoverImage:         product.CoverImage,
		TopicID:            product.TopicID,
		FolderID:           product.FolderID,
		QRCode:             product.QRCode,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          time.Now(),
	}

	err = s.productRepostitory.UpdateProduct(ctx, productData)
	if err != nil {
		return err
	}

	return nil

}

func (s *productService) DeleteProduct(ctx context.Context, id string) error {

	idObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.productRepostitory.DeleteProduct(ctx, idObjectID)

}
