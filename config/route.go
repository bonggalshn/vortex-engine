package config

import (
	"net/http"
	"vortex-engine/handler"

	"github.com/rs/cors"
)

func SetupRoutes() http.Handler {
	mainHandler := handler.ProcessMainHandler

	mux := http.NewServeMux()
	mux.HandleFunc("/main", mainHandler) // Main route

	// cors
	crs := allowCors()

	handlerWithCors := crs.Handler(mux)

	return handlerWithCors
}

func allowCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: GetAllowedOriginList(),
		AllowedMethods: GetAllowedMethodList(),
		AllowedHeaders: []string{"Content-Type", "Authorization", "Origin"},
	})
}
