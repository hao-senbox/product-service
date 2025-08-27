package folder

import "go.mongodb.org/mongo-driver/bson/primitive"

type Folder struct {
	ID       primitive.ObjectID  `json:"id" bson:"_id"`
	Name     string              `json:"name" bson:"name"`
	ParentID *primitive.ObjectID `json:"parent_id" bson:"parent_id"`
}
