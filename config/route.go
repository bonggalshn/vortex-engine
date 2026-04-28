package config

import (
	"net/http"
	"vortex-engine/handler"
)

func SetupRoutes() {
	// Main route
	http.HandleFunc("/main", handler.ProcessMainHandler)
}
