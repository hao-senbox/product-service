package folder

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, folderHandler *FolderHandler) {
	folderGroup := r.Group("api/v1/folders")
	{
		folderGroup.GET("", folderHandler.GetAllFolders)
		folderGroup.GET("/:id", folderHandler.GetFolder)
		folderGroup.POST("", folderHandler.CreateFolder)
		folderGroup.PUT("/:id", folderHandler.UpdateFolder)
		folderGroup.DELETE("/:id", folderHandler.DeleteFolder)
	}
}