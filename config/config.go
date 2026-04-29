package config

import (
	"os"
	"vortex-engine/logger"

	"github.com/joho/godotenv"
)

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
