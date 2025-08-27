package folder

import (
	"errors"
	"net/http"
	"product-service/helper"

	"github.com/gin-gonic/gin"
)

type FolderHandler struct {
	folderService FolderService
}

func NewFolderHandler(folderService FolderService) *FolderHandler {
	return &FolderHandler{
		folderService: folderService,
	}
}

func (h *FolderHandler) CreateFolder(ctx *gin.Context) {

	var req CreateFolderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.SendError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	res, err := h.folderService.CreateFolder(ctx, &req)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Folder created successfully", res)
}

func (h *FolderHandler) GetAllFolders(ctx *gin.Context) {

	res, err := h.folderService.GetAllFolders(ctx)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Folders retrieved successfully", res)

}

func (h *FolderHandler) GetFolder(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		helper.SendError(ctx, http.StatusBadRequest, errors.New("id is required"), nil)
		return
	}

	res, err := h.folderService.GetFolder(ctx, id)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Folder retrieved successfully", res)
}

func (h *FolderHandler) UpdateFolder(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		helper.SendError(ctx, http.StatusBadRequest, errors.New("id is required"), nil)
		return
	}

	var req UpdateFolderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.SendError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	err := h.folderService.UpdateFolder(ctx, &req, id)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Folder updated successfully", nil)
}

func (h *FolderHandler) DeleteFolder(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		helper.SendError(ctx, http.StatusBadRequest, errors.New("id is required"), nil)
		return
	}

	err := h.folderService.DeleteFolder(ctx, id)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Folder deleted successfully", nil)

}