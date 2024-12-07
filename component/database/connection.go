package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectToAzureSQL() (*gorm.DB, error) {
	// Try loading .env to read the Environment Variables.
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	// Get Connection String from Environment Variable
	connString := os.Getenv("DB_CONNECTION_STRING")
	if connString == "" {
		return nil, fmt.Errorf("DB_CONNECTION_STRING is not set in environment or .env file")
	}

	// Connect database with GORM
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Connected to Azure SQL successfully!")
	return db, nil
}
