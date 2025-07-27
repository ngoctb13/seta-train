package server

import (
	"github.com/gin-contrib/cors"
	hdl "github.com/ngoctb13/seta-train/rest-service/handler"
	"github.com/ngoctb13/seta-train/rest-service/internal/auth"
	team_usecases "github.com/ngoctb13/seta-train/rest-service/internal/domains/team/usecases"
	user_usecases "github.com/ngoctb13/seta-train/rest-service/internal/domains/user/usecases"
	"github.com/ngoctb13/seta-train/rest-service/repos"
)

type domains struct {
	user *user_usecases.User
	team *team_usecases.Team
}

func (s *Server) initCORS() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"*",
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"X-Access-Token",
		"X-Google-Access-Token",
	}
	s.router.Use(cors.New(corsConfig))
}

func (s *Server) initDomains(repo repos.IRepo) *domains {
	user := user_usecases.NewUser(repo.Users())
	team := team_usecases.NewTeam(repo.Teams())
	return &domains{
		user: user,
		team: team,
	}
}

func (s *Server) initRestRoute(domains *domains) {
	handler := hdl.NewHandler(domains.user, domains.team)

	routerAuth := s.router.Group("v1")
	routerAuth.Use(auth.AuthMiddleware())
	handler.ConfigAuthRouteAPI(routerAuth)
}

func (s *Server) initRouter(domains *domains) {
	//init rest route
	s.initRestRoute(domains)
}
