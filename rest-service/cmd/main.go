package main

import (
	"flag"

	"github.com/ngoctb13/seta-train/rest-service/server"
	"github.com/ngoctb13/seta-train/shared-modules/config"
	"github.com/ngoctb13/seta-train/shared-modules/setting"
	"go.uber.org/zap"
)

func main() {
	var configFile, port string
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	flag.StringVar(&port, "port", "", "Specify port")
	flag.Parse()

	defer setting.WaitOSSignal()

	//load config
	cfg, err := config.Load(configFile)
	if err != nil {
		zap.S().Errorf("load config fail with err: %v", err)
		panic(err)
	}

	// connect to db
	go setting.ConnectDatabase(cfg.DB)

	//start new server
	s := server.NewServer(cfg)
	s.Init()

	if err := s.ListenHTTP(); err != nil {
		zap.S().Errorf("start server fail with err: %v", err)
		panic(err)
	}
}
