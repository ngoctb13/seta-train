package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ngoctb13/seta-train/auth-service/server"
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
	logger := logger.InitLogger("auth-service")

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

	serverErr := make(chan error, 1)
	go func() {
		if err := s.ListenHTTP(); err != nil {
			serverErr <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stop:
		logger.Info("Received shutdown signal")
	case err := <-serverErr:
		logger.Error("Server error: %v", err)
	}

	//Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		logger.Error("Failed to shutdown server: %v", err)
	} else {
		logger.Info("Server shutdown completed")
	}
}
