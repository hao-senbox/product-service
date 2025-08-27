package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	ProductName        string             `json:"product_name" bson:"product_name"`
	OriginPriceStore   float64            `json:"original_price_store" bson:"original_price_store"`
	OriginPriceService float64            `json:"original_price_service" bson:"original_price_service"`
	ProductDescription string             `json:"product_description" bson:"product_description"`
	CoverImage         string             `json:"cover_image" bson:"cover_image"`
	TopicID            primitive.ObjectID `json:"topic_id" bson:"topic_id"`
	FolderID           primitive.ObjectID `json:"folder_id" bson:"folder_id"`
	QRCode             string             `json:"qrcode" bson:"qrcode"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at" bson:"updated_at"`
}
