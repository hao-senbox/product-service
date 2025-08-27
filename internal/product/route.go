package product

import (
	"product-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, ProductHandler *ProductHandler) {
	productGroup := r.Group("api/v1/products", middleware.Secured())
	{
		productGroup.GET("", ProductHandler.GetAllProducts)
		productGroup.GET("/:id", ProductHandler.GetProduct)
		productGroup.POST("", ProductHandler.CreateProduct)
		productGroup.PUT("/:id", ProductHandler.UpdateProduct)
		productGroup.DELETE("/:id", ProductHandler.DeleteProduct)
	}
}