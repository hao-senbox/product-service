package ports

import (
	"context"
	"product-service/internal/folder"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FolderRepository interface {
	GetFolder(ctx context.Context, id primitive.ObjectID) (*folder.Folder, error)
}