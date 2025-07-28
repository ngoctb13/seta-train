package server

import (
	"github.com/gin-contrib/cors"
	hdl "github.com/ngoctb13/seta-train/rest-service/handler"
	"github.com/ngoctb13/seta-train/rest-service/internal/auth"
	team_usecases "github.com/ngoctb13/seta-train/rest-service/internal/domains/team/usecases"
	"github.com/ngoctb13/seta-train/rest-service/repos"
	"github.com/ngoctb13/seta-train/shared-modules/infra/transaction"
)

type domains struct {
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

func (s *Server) initDomains(repo repos.IRepo, txn transaction.TxnManager) *domains {
	team := team_usecases.NewTeam(repo.Teams(), txn)
	return &domains{
		team: team,
	}
}

func (s *Server) initRestRoute(domains *domains) {
	handler := hdl.NewHandler(domains.team)

	routerAuth := s.router.Group("v1")
	routerAuth.Use(auth.AuthMiddleware())
	handler.ConfigAuthRouteAPI(routerAuth)
}

func (s *Server) initRouter(domains *domains) {
	//init rest route
	s.initRestRoute(domains)
}
