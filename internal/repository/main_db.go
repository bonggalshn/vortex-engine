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

	connectionString := "postgresql://vortex_user:sGrYCfMBcWnRRXIPKUehDeiORkINFZVU@dpg-d7og4g0sfn5c739chsrg-a.singapore-postgres.render.com/vortex_data?sslmode=require"

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
