package server

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/seta-train/rest-service/repos"
	"github.com/ngoctb13/seta-train/shared-modules/config"
	"github.com/ngoctb13/seta-train/shared-modules/infra"
	"github.com/ngoctb13/seta-train/shared-modules/infra/transaction"
	"github.com/ngoctb13/seta-train/shared-modules/utils"
	"go.uber.org/zap"
)

type Server struct {
	httpServer *http.Server
	router     *gin.Engine
	cfg        *config.AppConfig
	logger     *utils.Logger
}

func NewServer(cfg *config.AppConfig) *Server {
	router := gin.New()
	logger := utils.NewLogger("rest-service")
	return &Server{
		router: router,
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) Init() {
	db, err := infra.InitPostgres(s.cfg.DB)
	if err != nil {
		s.logger.Error("Failed to initialize database: %v", err)
		zap.S().Errorf("Init db error: %v", err)
		panic(err)
	}

	txnHelper := transaction.NewGormTxnManager(db)
	repo := repos.NewSQLRepo(db, s.cfg.DB)
	domains := s.initDomains(repo, txnHelper)
	s.initCORS()
	s.initRouter(domains)
}

// ListenHTTP ...
func (s *Server) ListenHTTP() error {
	listen, err := net.Listen("tcp", ":"+os.Getenv("REST_PORT"))
	if err != nil {
		s.logger.Error("Failed to create listener: %v", err)
		zap.S().Errorf("err %v", err)
		panic(err)
	}
	address := fmt.Sprintf(":%s", os.Getenv("REST_PORT"))
	fmt.Println(address)
	s.httpServer = &http.Server{
		Handler: s.router,
		Addr:    address,
	}

	s.logger.Info("Starting HTTP server on port %s", os.Getenv("REST_PORT"))
	zap.S().Infof("starting http server at port %v ...", os.Getenv("REST_PORT"))

	return s.httpServer.Serve(listen)
}
