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
			h.logger.Error("Unauthorized team creation attempt by user %s with role %v", userID, role)
			c.JSON(http.StatusForbidden, gin.H{"error": "only manager can create team"})
			return
		}

		var input models.CreateTeamReqeust
		if err := c.ShouldBindJSON(&input); err != nil {
			h.logger.Error("Failed to bind JSON for team creation: %v", err)
			log.Printf("binding json error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.team.CreateTeam(c, &useCaseModel.CreateTeamInput{
			TeamName: input.TeamName,
			UserID:   userID.(string),
		})

		if err != nil {
			h.logger.Error("Failed to create team '%s': %v", input.TeamName, err)
			log.Printf("CreateTeamUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		h.logger.Info("Team '%s' created successfully by user %s", input.TeamName, userID)
		c.JSON(http.StatusOK, gin.H{"message": "Team created successfully"})
	}
}

func (h *Handler) AddTeamMembersHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get(userRoleKey)
		userID, _ := c.Get(userIDKey)
		if !ok || role != managerRole {
			h.logger.Error("Unauthorized team member addition attempt by user %s with role %v", userID, role)
			c.JSON(http.StatusForbidden, gin.H{"error": "only manager can do"})
			return
		}

		teamID := c.Param("teamId")
		if teamID == "" {
			h.logger.Error("Team ID is missing in request")
			c.JSON(http.StatusBadRequest, gin.H{"error": "team ID is required"})
			return
		}

		var input models.AddTeamMembersRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			h.logger.Error("Failed to bind JSON for adding team members: %v", err)
			log.Printf("binding json error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.team.AddTeamMembers(c, &useCaseModel.AddTeamMembersInput{
			TeamID:    teamID,
			UserIDs:   input.UserIDs,
			CurUserID: userID.(string),
		})

		if err != nil {
			h.logger.Error("Failed to add members to team %s: %v", teamID, err)
			log.Printf("AddTeamMembers usecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		h.logger.Info("Successfully added %d members to team %s", len(input.UserIDs), teamID)
		c.JSON(http.StatusOK, gin.H{"message": "Members added successfully"})
	}
}

func (h *Handler) AddTeamManagersHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get(userRoleKey)
		userID, _ := c.Get(userIDKey)
		if !ok || role != managerRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "only manager can do"})
			return
		}

		teamID := c.Param("teamId")
		if teamID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "team ID is required"})
			return
		}

		var input models.AddTeamManagersRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			log.Printf("binding json error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.team.AddTeamManagers(c, &useCaseModel.AddTeamManagersInput{
			TeamID:    teamID,
			UserIDs:   input.UserIDs,
			CurUserID: userID.(string),
		})

		if err != nil {
			log.Printf("AddTeamManagers usecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Managers added successfully"})
	}
}

func (h *Handler) RemoveTeamMemberHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get(userRoleKey)
		userID, _ := c.Get(userIDKey)
		if !ok || role != managerRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "only manager can do"})
			return
		}

		teamID := c.Param("teamId")
		if teamID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "team ID is required"})
			return
		}

		memberID := c.Param("memberId")
		if memberID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "member ID is required"})
			return
		}

		err := h.team.RemoveTeamMember(c, &useCaseModel.RemoveTeamMemberInput{
			TeamID:    teamID,
			MemberID:  memberID,
			CurUserID: userID.(string),
		})

		if err != nil {
			log.Printf("RemoveTeamMember usecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Member removed successfully"})
	}
}

func (h *Handler) RemoveTeamManagerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get(userRoleKey)
		userID, _ := c.Get(userIDKey)
		if !ok || role != managerRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "only manager can create team"})
			return
		}

		teamID := c.Param("teamId")
		if teamID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "team ID is required"})
			return
		}

		managerID := c.Param("managerId")
		if managerID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "manager ID is required"})
			return
		}

		err := h.team.RemoveTeamManager(c, &useCaseModel.RemoveTeamManagerInput{
			CurUserID: userID.(string),
			TeamID:    teamID,
			ManagerID: managerID,
		})

		if err != nil {
			log.Printf("RemoveTeamManager usecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Manager removed successfully"})
	}
}
