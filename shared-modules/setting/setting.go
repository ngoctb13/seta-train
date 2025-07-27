package setting

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ngoctb13/seta-train/shared-modules/config"
	"github.com/ngoctb13/seta-train/shared-modules/infra"
	"go.uber.org/zap"
)

const (
	migrationFile = "file://./migrations/sql"
)

// Connect to database...
func ConnectDatabase(cfg *config.PostgresConfig) {
	infra.CreateDB(cfg, migrationFile)
}

// WaitOSSignal...
func WaitOSSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	s := <-c
	zap.S().Infof("Receive os.Signal: %s", s.String())
}
