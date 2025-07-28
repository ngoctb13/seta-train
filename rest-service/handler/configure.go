package handler

import (
	"github.com/gin-gonic/gin"
	team_usecases "github.com/ngoctb13/seta-train/rest-service/internal/domains/team/usecases"
)

type Handler struct {
	team *team_usecases.Team
}

func NewHandler(team *team_usecases.Team) *Handler {
	return &Handler{
		team: team,
	}
}

func (h *Handler) ConfigAuthRouteAPI(router *gin.RouterGroup) {
	router.GET("/hello", h.HelloHandler())

	// team routes
	router.POST("/teams", h.CreateTeamHandler())
	router.POST("/teams/:teamId/members", h.AddTeamMembersHandler())
	router.POST("/teams/:teamId/managers", h.AddTeamManagersHandler())
	router.DELETE("/teams/:teamId/members/:memberId", h.RemoveTeamMemberHandler())
	router.DELETE("/teams/:teamId/managers/:managerId", h.RemoveTeamManagerHandler())
}
