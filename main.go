package main

import (
	"fmt"
	"net/http"
	"vortex-engine/config"
)

func main() {
	fmt.Println("== Starting Vortex Engine ==")

	// Setup http routes
	config.SetupRoutes()
	// Setup logging
	config.SetupLogging()

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)

}
