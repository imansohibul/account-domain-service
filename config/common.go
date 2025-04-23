package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type ServiceConfig struct {
	DatabaseConfig DatabaseConfig `envconfig:"DB"`
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() (ServiceConfig, error) {
	var cfg ServiceConfig

	// load from .env if exists
	if _, err := os.Stat(".env"); err == nil {
		if err := gotenv.Load(); err != nil {
			return cfg, err
		}
	}

	// parse environment variable to config struct
	err := envconfig.Process("service", &cfg)
	return cfg, err
}

type DatabaseConfig struct {
	Host     string `envconfig:"HOST"`
	Port     int    `envconfig:"PORT"`
	Username string `envconfig:"USERNAME"`
	Password string `envconfig:"PASSWORD"`
	Database string `envconfig:"NAME"`
}

// BuildDSN constructs the PostgreSQL DSN in URL format
func (db DatabaseConfig) PostgresDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Database,
	)
}

func initPostgresDatabase(cfg ServiceConfig) rel.Repository {
	dsn := cfg.DatabaseConfig.PostgresDSN()

	log.Println("Connecting to PostgreSQL database...", dsn)
	adapter, err := postgres.Open(cfg.DatabaseConfig.PostgresDSN())
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	if err := adapter.Ping(context.Background()); err != nil {
		adapter.Close()
		panic(fmt.Sprintf("failed to ping database: %v", err))
	}

	return rel.New(adapter)
}
