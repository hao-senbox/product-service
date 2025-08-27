package product

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductResponse struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	ProductName        string             `json:"product_name" bson:"product_name"`
	OriginPriceStore   float64            `json:"original_price_store" bson:"original_price_store"`
	OriginPriceService float64            `json:"original_price_service" bson:"original_price_service"`
	ProductDescription string             `json:"product_description" bson:"product_description"`
	CoverImage         string             `json:"cover_image" bson:"cover_image"`
	Topic              *Topic              `json:"topic" bson:"topic"`
	Folder             Folder             `json:"folder" bson:"folder"`
	QRCode             string             `json:"qrcode" bson:"qrcode"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at" bson:"updated_at"`
}

type Topic struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type Folder struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
