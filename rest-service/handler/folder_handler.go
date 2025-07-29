package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/seta-train/rest-service/handler/models"
	useCaseModel "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
)

func (h *Handler) CreateFolderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CreateFolderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("binding json error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID, _ := c.Get(userIDKey)

		err := h.folder.CreateFolder(c, &useCaseModel.CreateFolderInput{
			FolderName: req.FolderName,
			UserID:     userID.(string),
		})

		if err != nil {
			log.Printf("CreateFolderUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Folder created successfully"})
	}
}
