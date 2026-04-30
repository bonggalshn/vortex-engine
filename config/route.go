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
	c := cors.New(cors.Options{
		AllowedOrigins: GetAllowedOriginList(),
	})

	handlerWithCors := c.Handler(mux)

	return handlerWithCors
}

func allowCors() {

}
