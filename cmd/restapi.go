package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
	"imansohibul.my.id/account-domain-service/config"
	"imansohibul.my.id/account-domain-service/internal/rest/server"
	"imansohibul.my.id/account-domain-service/util"
)

var logger util.Logger

func init() {
	logger = util.GetZapLogger()
}

func RestAPI(c *cli.Context) error {
	address := c.String("address")
	ctx := context.Background()

	restAPIServer, err := config.NewRestAPI()
	if err != nil {
		logger.Fatal(ctx, "failed to initialize REST API server", err, nil)
	}

	// Graceful shutdown handler
	idleConnsClosed := make(chan struct{})
	go handleGracefulShutdown(ctx, restAPIServer, idleConnsClosed)

	logger.Info(ctx, "Starting REST API server...", nil)
	if err := restAPIServer.Start(address); err != http.ErrServerClosed {
		logger.Fatal(ctx, "REST API server stopped with error", err, nil)
	}

	<-idleConnsClosed
	logger.Info(ctx, "Server shut down gracefully", nil)
	return nil
}

func handleGracefulShutdown(ctx context.Context, restAPIServer *server.RestAPIServer, done chan struct{}) {
	// Listen for interrupt signal (e.g., Ctrl+C, SIGTERM)
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	logger.Warn(ctx, "Shutdown signal received", nil)

	if err := restAPIServer.Shutdown(ctx); err != nil {
		logger.Fatal(ctx, "Error during server shutdown", err, nil)
	}

	close(done)
}
