package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
	"imansohibul.my.id/account-domain-service/util"
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

func initPostgresDatabase(cfg ServiceConfig) (rel.Repository, error) {
	adapter, err := postgres.Open(cfg.DatabaseConfig.PostgresDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	if err := adapter.Ping(context.Background()); err != nil {
		adapter.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	rel := rel.New(adapter)
	rel.Instrumentation(DatabaseLogger)

	return rel, nil
}

// DatabaseLogger instrumentation to log queries and rel operation.
func DatabaseLogger(ctx context.Context, op string, message string, args ...any) func(err error) {
	start := time.Now()

	return func(err error) {
		duration := time.Since(start)
		fields := map[string]interface{}{
			"op":       op,
			"duration": duration,
		}

		if err != nil {
			util.GetZapLogger().Error(ctx, message, err, fields)
		} else {
			util.GetZapLogger().Info(ctx, message, fields)
		}
	}
}
