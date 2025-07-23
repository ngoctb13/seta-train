package main

import (
	"flag"

	"github.com/ngoctb13/seta-train/config"
	"github.com/ngoctb13/seta-train/server"
	"github.com/ngoctb13/seta-train/setting"
	"go.uber.org/zap"
)

// Optional: put code in src folder
// The outside contain other like config, deploy, migrations, etc.
func main() {
	var configFile, port string
	// Todo should add a default value for config file
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	// put in config file
	flag.StringVar(&port, "port", "", "Specify port")
	flag.Parse()

	defer setting.WaitOSSignal()

	//load config
	cfg, err := config.Load(configFile)
	if err != nil {
		zap.S().Errorf("load config fail with err: %v", err)
		panic(err)
	}

	// Todo: should add config migratedb: bool in config file
	// not every time we run the server, we need to migrate db
	go setting.MigrateDatabase(cfg.DB)

	//start new server
	s := server.NewServer(cfg)
	s.Init()

	// add graceful shutdown base on os signal
	if err := s.ListenHTTP(); err != nil {
		zap.S().Errorf("start server fail with err: %v", err)
		panic(err)
	}
}
