package product

type CreateProductRequest struct {
	ProductName        string  `json:"product_name" bson:"product_name"`
	OriginPriceStore   float64 `json:"original_price_store" bson:"original_price_store"`
	OriginPriceService float64 `json:"original_price_service" bson:"original_price_service"`
	ProductDescription string  `json:"product_description" bson:"product_description"`
	CoverImage         string  `json:"cover_image" bson:"cover_image"`
	TopicID            string  `json:"topic_id" bson:"topic_id"`
	FolderID           string  `json:"folder_id" bson:"folder_id"`
	QRCode             string  `json:"qrcode" bson:"qrcode"`
}

type UpdateProductRequest struct {
	ProductName        string  `json:"product_name" bson:"product_name"`
	OriginPriceStore   float64 `json:"original_price_store" bson:"original_price_store"`
	OriginPriceService float64 `json:"original_price_service" bson:"original_price_service"`
	ProductDescription string  `json:"product_description" bson:"product_description"`
	CoverImage         string  `json:"cover_image" bson:"cover_image"`
	TopicID            string  `json:"topic_id" bson:"topic_id"`
	FolderID           string  `json:"folder_id" bson:"folder_id"`
	QRCode             string  `json:"qrcode" bson:"qrcode"`
}
