package main

import (
	"flag"

	"github.com/ngoctb13/seta-train/rest-service/server"
	"github.com/ngoctb13/seta-train/shared-modules/config"
	"github.com/ngoctb13/seta-train/shared-modules/logger"
	"github.com/ngoctb13/seta-train/shared-modules/setting"
)

func main() {
	var configFile, port string
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	flag.StringVar(&port, "port", "", "Specify port")
	flag.Parse()

	defer setting.WaitOSSignal()

	// Initialize logger
	logger := logger.InitLogger("rest-service")

	//load config
	cfg, err := config.Load(configFile)
	if err != nil {
		logger.Error("Failed to load config: %v", err)
		panic(err)
	}

	// connect to db
	go setting.ConnectDatabase(cfg.DB)

	//start new server
	s := server.NewServer(cfg, logger)
	s.Init()

	if err := s.ListenHTTP(); err != nil {
		logger.Error("Failed to start server: %v", err)
		panic(err)
	}
}
