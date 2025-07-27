package infra

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/cenkalti/backoff"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ngoctb13/seta-train/shared-modules/config"
)

func InitPostgres(cfg *config.PostgresConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("get DB instance failed: %v", err)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifeTimeMiliseconds) * time.Millisecond)

	return db, nil
}

func CreateDB(cfg *config.PostgresConfig, migrationFile string) *gorm.DB {
	var db *gorm.DB
	boff := backoff.NewExponentialBackOff()

	err := backoff.Retry(func() error {
		var errNested error
		db, errNested = InitPostgres(cfg)
		if errNested != nil {
			fmt.Printf("Connect postgres error: %s\n", errNested.Error())
		} else {
			fmt.Println("Connect postgres successful.")
		}
		return errNested
	}, boff)

	if err != nil {
		panic(err)
	}

	return db
}
