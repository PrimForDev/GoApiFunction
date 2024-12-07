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
	// ลองโหลด .env เพื่ออ่านค่า Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	// รับ Connection String จาก Environment Variable
	connString := os.Getenv("DB_CONNECTION_STRING")
	if connString == "" {
		return nil, fmt.Errorf("DB_CONNECTION_STRING is not set in environment or .env file")
	}

	// เชื่อมต่อฐานข้อมูลด้วย GORM
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Connected to Azure SQL successfully!")
	return db, nil
}
