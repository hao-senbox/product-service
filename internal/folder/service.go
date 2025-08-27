package folder

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FolderService interface {
	CreateFolder(ctx context.Context, req *CreateFolderRequest) (string, error)
	GetAllFolders(ctx context.Context) ([]*Folder, error)
	GetFolder(ctx context.Context, id string) (*Folder, error)
	UpdateFolder(ctx context.Context, req *UpdateFolderRequest, id string) error
	DeleteFolder(ctx context.Context, id string) error
}

type folderService struct {
	folderReposity FolderRepository
}

func NewFolderService(folderRepository FolderRepository) FolderService {
	return &folderService{
		folderReposity: folderRepository,
	}
}

func (s *folderService) CreateFolder(ctx context.Context, req *CreateFolderRequest) (string, error) {

	var parentID *primitive.ObjectID

	if req.Name == "" {
		return "", errors.New("name is required")
	}

	if req.ParentID != nil {
		result, err := primitive.ObjectIDFromHex(*req.ParentID)
		if err != nil {
			return "", err
		}

		parentID = &result
	} else {
		parentID = nil
	}

	folder := &Folder{
		ID:       primitive.NewObjectID(),
		Name:     req.Name,
		ParentID: parentID,
	}

	return s.folderReposity.CreateFolder(ctx, folder)

}

func (s *folderService) GetAllFolders(ctx context.Context) ([]*Folder, error) {
	return s.folderReposity.GetAllFolders(ctx)
}

func (s *folderService) GetFolder(ctx context.Context, id string) (*Folder, error) {

	if id == "" {
		return nil, errors.New("id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return s.folderReposity.GetFolder(ctx, objectID)
}

func (s *folderService) UpdateFolder(ctx context.Context, req *UpdateFolderRequest, id string) error {

	var parentID *primitive.ObjectID

	if id == "" {
		return errors.New("id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	folder, err := s.folderReposity.GetFolder(ctx, objectID)
	if err != nil {
		return err
	}

	if req.Name != "" {
		folder.Name = req.Name
	}

	if req.ParentID != nil {
		result, err := primitive.ObjectIDFromHex(*req.ParentID)
		if err != nil {
			return err
		}
		parentID = &result
		folder.ParentID = parentID
	}

	err = s.folderReposity.UpdateFolder(ctx, folder)
	if err != nil {
		return err
	}

	return nil
}

func (s *folderService) DeleteFolder(ctx context.Context, id string) error {

	if id == "" {
		return errors.New("id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.folderReposity.DeleteFolder(ctx, objectID)

}
