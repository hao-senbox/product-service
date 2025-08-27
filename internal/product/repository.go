package product

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *Product) (string, error)
	GetAllProducts(ctx context.Context) ([]*Product, error)
	GetProduct(ctx context.Context, id primitive.ObjectID) (*Product, error)
	UpdateProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, id primitive.ObjectID) error
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) ProductRepository {
	return &productRepository{
		collection: collection,
	}
}

func (r *productRepository) CreateProduct(ctx context.Context, product *Product) (string, error) {
	
	result, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return "", err
	}
	
	return result.InsertedID.(primitive.ObjectID).Hex(), nil

}

func (r *productRepository) GetAllProducts(ctx context.Context) ([]*Product, error) {

	var products []*Product

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
	
}

func (r *productRepository) GetProduct(ctx context.Context, id primitive.ObjectID) (*Product, error) {

	var product Product

	filter := bson.M{"_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil

}

func (r *productRepository) UpdateProduct(ctx context.Context, product *Product) error {

	filter := bson.M{"_id": product.ID}

	update := bson.M{"$set": product}
	
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	
	return nil

}

func (r *productRepository) DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
	
	filter := bson.M{"_id": id}
	
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	
	return nil
	
}