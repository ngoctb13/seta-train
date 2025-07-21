package infra

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/cenkalti/backoff"
	"github.com/golang-migrate/migrate/v4"
	migratePg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ngoctb13/seta-train/config"
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

func CreateDBAndMigrate(cfg *config.PostgresConfig, migrationFile string) *gorm.DB {
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

	Migrate(migrationFile, cfg.MigrationConnURL)
	return db
}

func Migrate(source string, connStr string) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println("Migrating....")
	fmt.Printf("Source=%+v Connection=%+v\n", source, connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	driver, err := migratePg.WithInstance(db, &migratePg.Config{})
	if err != nil {
		panic(err)
	}

	mg, err := migrate.NewWithDatabaseInstance(
		source,
		"postgres",
		driver,
	)
	if err != nil {
		panic(err)
	}
	defer mg.Close()

	version, dirty, err := mg.Version()
	if err != nil && err.Error() != migrate.ErrNilVersion.Error() {
		panic(err)
	}

	if dirty {
		_ = mg.Force(int(version) - 1) // force to clean state
	}

	err = mg.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	fmt.Println("Migration done...")
}
