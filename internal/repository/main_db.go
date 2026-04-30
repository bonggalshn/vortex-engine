package repository

import (
	"database/sql"
	"os"
	"vortex-engine/logger"

	_ "github.com/lib/pq"
)

func EstablishMainDBConnection() {
	logger.Info.Println("Establish main database connection.")

	dbHost := os.Getenv("DATABASE_HOST")
	isDatabaseHostSet := dbHost != ""
	logger.Info.Printf("Database host: %t", isDatabaseHostSet)

	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	connectionString := "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + "/" + dbName + "?sslmode=require"

	dbConnection, dbError := sql.Open("postgres", connectionString)
	if dbError != nil {
		logger.Error.Println("Error establishing main database connection:", dbError)
		return
	}

	if dbError := dbConnection.Ping(); dbError != nil {
		logger.Error.Println("Error pinging main database:", dbError)
		return
	}

	logger.Info.Println("Main database connection established successfully.")
}
