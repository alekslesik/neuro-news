// cmd/neuro-news/main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"githhub.com/alekslesik/neuro-news/internal/app/handler"
	"githhub.com/alekslesik/neuro-news/internal/config"
	"githhub.com/alekslesik/neuro-news/internal/infra/logger"
	"githhub.com/alekslesik/neuro-news/internal/web/router"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Initialize logger
	log, err := logger.New(cfg.LogLevel)
	if err != nil {
		log.Fatal("Error initializing logger:", err)
	}

	// Create HTTP router
	r := router.New()

	// Set up HTTP routes
	handler.SetupRoutes(r)

	// Start HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Printf("Starting Neuro News application on :%d...\n", cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
