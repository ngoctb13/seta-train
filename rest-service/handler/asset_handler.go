package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	useCaseModel "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
)

func (h *Handler) GetAssetsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		teamID := c.Param("teamId")
		if teamID == "" {
			h.logger.Error("team ID is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "team ID is required"})
		}

		userID, _ := c.Get(userIDKey)

		assets, err := h.asset.GetAssets(c, &useCaseModel.GetAssetsInput{
			TeamID: teamID,
			UserID: userID.(string),
		})

		if err != nil {
			h.logger.Error("GetAssets fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		h.logger.Info("GetAssetsHandler success")
		c.JSON(http.StatusOK, gin.H{"assets": assets})
	}
}
