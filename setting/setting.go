package setting

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ngoctb13/seta-train/config"
	"github.com/ngoctb13/seta-train/infra"
	"go.uber.org/zap"
)

const (
	migrationFile = "file://./migrations/sql"
)

// Migrate database...
func MigrateDatabase(cfg *config.PostgresConfig) {
	infra.CreateDBAndMigrate(cfg, migrationFile)
}

// WaitOSSignal...
func WaitOSSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	s := <-c
	zap.S().Infof("Receive os.Signal: %s", s.String())
}
