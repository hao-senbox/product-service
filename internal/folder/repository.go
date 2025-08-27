package folder

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FolderRepository interface {
	CreateFolder(ctx context.Context, folder *Folder) (string, error)
	GetAllFolders(ctx context.Context) ([]*Folder, error)
	GetFolder(ctx context.Context, id primitive.ObjectID) (*Folder, error)
	UpdateFolder(ctx context.Context, folder *Folder) error
	DeleteFolder(ctx context.Context, id primitive.ObjectID) error
}

type folderRepository struct {
	collection *mongo.Collection
}

func NewFolderRepository(collection *mongo.Collection) FolderRepository {
	return &folderRepository{
		collection: collection,
	}
}

func (r *folderRepository) CreateFolder(ctx context.Context, folder *Folder) (string, error) {
	
	result, err := r.collection.InsertOne(ctx, folder)
	if err != nil {
		return "", err
	}
	
	return result.InsertedID.(primitive.ObjectID).Hex(), nil

}

func (r *folderRepository) GetAllFolders(ctx context.Context) ([]*Folder, error) {

	var folders []*Folder

	filer := bson.M{}

	cursor, err := r.collection.Find(ctx, filer)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &folders)
	if err != nil {
		return nil, err
	}

	return folders, nil
}

func (r *folderRepository) GetFolder(ctx context.Context, id primitive.ObjectID) (*Folder, error) {

	var folder Folder

	filter := bson.M{"_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&folder)
	if err != nil {
		return nil, err
	}

	return &folder, nil

}

func (r *folderRepository) UpdateFolder(ctx context.Context, folder *Folder) error {
	
	filter := bson.M{"_id": folder.ID}

	update := bson.M{"$set": folder}
	
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	
	return nil

}

func (r *folderRepository) DeleteFolder(ctx context.Context, id primitive.ObjectID) error {
	
	filter := bson.M{"_id": id}
	
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	
	return nil
}