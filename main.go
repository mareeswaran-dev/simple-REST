package main

import (
	"fmt"
	"log"

	"user-crud-api/app"
	"user-crud-api/config"
	"user-crud-api/routes"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize application
	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("Error initializing application: %v", err)
	}
	defer application.Close()

	// Setup and run the router
	r := routes.SetupRouter(application)
	fmt.Printf("Server is running on port %s\n", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
