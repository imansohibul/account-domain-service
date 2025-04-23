package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"imansohibul.my.id/account-domain-service/config"
	"imansohibul.my.id/account-domain-service/internal/rest/server"
)

func main() {
	ctx := context.Background()

	restAPIServer, err := config.NewRestAPI()
	if err != nil {
		log.Fatalf("failed to initialize REST API server: %v", err)
	}

	// Graceful shutdown handler
	idleConnsClosed := make(chan struct{})
	go handleGracefulShutdown(ctx, restAPIServer, idleConnsClosed)

	log.Println("Starting REST API server...")
	if err := restAPIServer.Start(); err != nil {
		log.Printf("REST API server stopped with error: %v", err)
	}

	<-idleConnsClosed
	log.Println("Server shut down gracefully")
}

func handleGracefulShutdown(ctx context.Context, restAPIServer *server.RestAPIServer, done chan struct{}) {
	// Listen for interrupt signal (e.g., Ctrl+C, SIGTERM)
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	log.Println("Shutdown signal received")

	if err := restAPIServer.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	}

	close(done)
}
