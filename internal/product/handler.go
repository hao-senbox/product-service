package product

import (
	"context"
	"errors"
	"net/http"
	"product-service/helper"
	"product-service/pkg/constants"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService ProductService
}

func NewProductHandler(ProductService ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: ProductService,
	}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {

	var req CreateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.SendError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	res, err := h.ProductService.CreateProduct(ctx, &req)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Product created successfully", res)

}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {

	token, ok := ctx.Get(constants.Token)
	if !ok {
		helper.SendError(ctx, http.StatusBadRequest, errors.New("token not found"), nil)
		return
	}

	c := context.WithValue(ctx, constants.TokenKey, token)

	res, err := h.ProductService.GetAllProducts(c)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Products retrieved successfully", res)

}

func (h *ProductHandler) GetProduct(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		helper.SendError(ctx, http.StatusBadRequest, errors.New("id is required"), nil)
		return
	}

	token, ok := ctx.Get(constants.Token)
	if !ok {
		helper.SendError(ctx, http.StatusBadRequest, errors.New("token not found"), nil)
		return
	}

	c := context.WithValue(ctx, constants.TokenKey, token)

	res, err := h.ProductService.GetProduct(c, id)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Product retrieved successfully", res)

}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		helper.SendError(ctx, http.StatusBadRequest, errors.New("id is required"), nil)
		return
	}

	var req UpdateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.SendError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	err := h.ProductService.UpdateProduct(ctx, &req, id)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Product updated successfully", nil)

}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		helper.SendError(ctx, http.StatusBadRequest, errors.New("id is required"), nil)
		return
	}

	err := h.ProductService.DeleteProduct(ctx, id)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Product deleted successfully", nil)

}
