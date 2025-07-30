package server

import (
	"github.com/gin-contrib/cors"
	hdl "github.com/ngoctb13/seta-train/rest-service/handler"
	"github.com/ngoctb13/seta-train/rest-service/internal/auth"
	folder_usecases "github.com/ngoctb13/seta-train/rest-service/internal/domains/folder/usecases"
	team_usecases "github.com/ngoctb13/seta-train/rest-service/internal/domains/team/usecases"
	"github.com/ngoctb13/seta-train/rest-service/repos"
	"github.com/ngoctb13/seta-train/shared-modules/infra/transaction"
	"github.com/ngoctb13/seta-train/shared-modules/logger"
)

type domains struct {
	team   *team_usecases.Team
	folder *folder_usecases.Folder
	note   *folder_usecases.Note
	asset  *folder_usecases.Asset
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
	folder := folder_usecases.NewFolder(repo.Folders(), repo.Notes(), txn)
	note := folder_usecases.NewNote(repo.Notes(), repo.Folders(), txn)
	asset := folder_usecases.NewAsset(repo.Folders(), repo.Notes(), repo.Teams(), txn)
	return &domains{
		team:   team,
		folder: folder,
		note:   note,
		asset:  asset,
	}
}

func (s *Server) initRestRoute(domains *domains) {
	handler := hdl.NewHandler(domains.team, domains.folder, domains.note, domains.asset, s.logger)

	// Add logging middleware
	loggingMiddleware := logger.NewLoggingMiddleware(s.logger, "rest-service")
	s.router.Use(loggingMiddleware.LoggingMiddleware())

	routerAuth := s.router.Group("v1")
	routerAuth.Use(auth.AuthMiddleware())
	handler.ConfigAuthRouteAPI(routerAuth)
}

func (s *Server) initRouter(domains *domains) {
	//init rest route
	s.initRestRoute(domains)
}
