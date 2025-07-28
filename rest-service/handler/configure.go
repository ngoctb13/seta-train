package handler

import (
	"github.com/gin-gonic/gin"
	team_usecases "github.com/ngoctb13/seta-train/rest-service/internal/domains/team/usecases"
	user_usecases "github.com/ngoctb13/seta-train/rest-service/internal/domains/user/usecases"
)

type Handler struct {
	user *user_usecases.User
	team *team_usecases.Team
}

func NewHandler(user *user_usecases.User, team *team_usecases.Team) *Handler {
	return &Handler{
		user: user,
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
