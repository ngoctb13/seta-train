package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ngoctb13/seta-train/rest-service/repos"
	"github.com/ngoctb13/seta-train/rest-service/worker"
	"github.com/ngoctb13/seta-train/shared-modules/config"
	"github.com/ngoctb13/seta-train/shared-modules/infra"
	"github.com/ngoctb13/seta-train/shared-modules/kafka"
	"github.com/ngoctb13/seta-train/shared-modules/setting"
	"go.uber.org/zap"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	flag.Parse()

	defer setting.WaitOSSignal()

	// load config
	cfg, err := config.Load(configFile)
	if err != nil {
		panic(err)
	}

	db, err := infra.InitPostgres(cfg.DB)
	if err != nil {
		zap.S().Errorf("Init db error: %v", err)
		panic(err)
	}

	// init db repo, kafka
	repo := repos.NewSQLRepo(db, cfg.DB)

	opts := []kafka.ProducerOption{
		kafka.ProducerWithAckMode(kafka.AckModeInSync),
		kafka.ProducerWithAutoCreateTopics(),
	}
	producer, err := kafka.NewSyncProducer(cfg, cfg.Kafka.Brokers, opts...)
	if err != nil {
		zap.S().Errorf("Init kafka producer error: %v", err)
	}

	// worker
	w := worker.InitWorker(cfg, repo.OutgoingEvents(), producer)
	ctx, cancel := context.WithCancel(context.Background())
	w.Start(ctx)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	fmt.Println("Received interrupt signal, canceling job...")
	cancel()
	time.Sleep(time.Second)
	fmt.Println("Shutting down job...")
}
