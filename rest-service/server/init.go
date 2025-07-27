package server

import (
	"github.com/gin-contrib/cors"
	hdl "github.com/ngoctb13/seta-train/rest-service/handler"
	"github.com/ngoctb13/seta-train/rest-service/internal/auth"
	"github.com/ngoctb13/seta-train/rest-service/internal/domains/user/usecases"
	"github.com/ngoctb13/seta-train/rest-service/repos"
)

type domains struct {
	user *usecases.User
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
	user := usecases.NewUser(repo.Users())
	return &domains{
		user: user,
	}
}

func (s *Server) initRestRoute(domains *domains) {
	handler := hdl.NewHandler(domains.user)

	routerAuth := s.router.Group("v1")
	routerAuth.Use(auth.AuthMiddleware())
	handler.ConfigAuthRouteAPI(routerAuth)
}

func (s *Server) initRouter(domains *domains) {
	//init rest route
	s.initRestRoute(domains)
}
