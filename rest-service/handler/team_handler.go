package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/seta-train/rest-service/handler/models"
	useCaseModel "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
)

const (
	userIDKey   = "userID"
	userRoleKey = "userRole"
	managerRole = "MANAGER"
)

func (h *Handler) CreateTeamHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get(userRoleKey)
		userID, _ := c.Get(userIDKey)
		if !ok || role != managerRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "only manager can create team"})
			return
		}

		var input models.CreateTeamReqeust
		if err := c.ShouldBindJSON(&input); err != nil {
			log.Printf("binding json error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.team.CreateTeam(c, &useCaseModel.CreateTeamInput{
			TeamName: input.TeamName,
			UserID:   userID.(string),
		})

		if err != nil {
			log.Printf("CreateTeamUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Team created successfully"})
	}
}
