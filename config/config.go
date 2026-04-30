package config

import (
	"os"
	"vortex-engine/logger"

	"github.com/joho/godotenv"
)

var allowedOriginList = []string{"https://vortex-ui-eight.vercel.app"}
var allowedMethodList = []string{"GET", "POST", "PUT", "DELETE"}

func LoadEnvConfig() {
	if os.Getenv("ENV") != "" {
		logger.Info.Printf("Running in %s environment, skipping .env loading.", os.Getenv("ENV"))
		return
	}

	err := godotenv.Load()
	if err != nil {
		logger.Error.Println("No .env file found")
	}

	logger.Info.Println("Environment variables loaded successfully.")

}

func GetAllowedOriginList() []string {
	return allowedOriginList
}

func GetAllowedMethodList() []string {
	return allowedMethodList
}
