package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dokk-dev/opus/internal/ai"
	"github.com/dokk-dev/opus/internal/api"
	"github.com/dokk-dev/opus/internal/config"
	"github.com/dokk-dev/opus/internal/gateway"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Ollama client
	ollamaClient := ai.NewOllamaClient(cfg.OllamaURL, cfg.OllamaModel)

	// Check if Ollama is available
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if ollamaClient.IsAvailable(ctx) {
		log.Printf("Connected to Ollama at %s (model: %s)", cfg.OllamaURL, cfg.OllamaModel)
	} else {
		log.Printf("Warning: Ollama not available at %s - AI features will fail", cfg.OllamaURL)
	}
	cancel()

	// Initialize AI router with department agents
	aiRouter := ai.NewRouter(ollamaClient)

	// Initialize WebSocket gateway
	gw := gateway.New(cfg)

	// Initialize HTTP API
	router := api.NewRouter(cfg, gw, aiRouter)

	server := &http.Server{
		Addr:         cfg.ServerAddr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Opus server starting on %s", cfg.ServerAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}
