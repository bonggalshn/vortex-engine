package main

import (
	"net/http"
	"vortex-engine/config"
	"vortex-engine/internal/repository"
	"vortex-engine/logger"
)

func main() {
	// Setup logging
	logger.SetupLogging()
	logger.Info.Println("== Vortex Engine Start ==")

	// Load environment variables
	config.LoadEnvConfig()

	// Establish main database connection
	repository.EstablishMainDBConnection()

	// Setup http routes
	config.SetupRoutes()

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)

}
