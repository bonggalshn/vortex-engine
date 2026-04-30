package config

import (
	"os"
	"strings"
	"vortex-engine/logger"
	"vortex-engine/util"

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

func GetAllowedMethodList() []string {
	return allowedMethodList
}

func IsAllowOrigin(origin string) bool {
	if util.ListContains(allowedOriginList, origin) {
		return true
	}

	return strings.HasPrefix(origin, "https://vortex") &&
		strings.HasSuffix(origin, ".vercel.app")
}
