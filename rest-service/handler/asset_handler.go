package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	useCaseModel "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
)

func (h *Handler) GetAssetsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		teamID := c.Param("teamId")
		if teamID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "team ID is required"})
		}

		userID, _ := c.Get(userIDKey)

		assets, err := h.asset.GetAssets(c, &useCaseModel.GetAssetsInput{
			TeamID: teamID,
			UserID: userID.(string),
		})

		if err != nil {
			log.Printf("GetAssetsUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"assets": assets})
	}
}
