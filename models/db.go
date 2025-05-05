package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection and runs migrations
func ConnectDB() {
	dbPath := os.Getenv("DBPATH")
	if dbPath == "" {
		log.Fatal("DBPATH environment variable not set")
	}

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = database

	// Auto migrate all models
	err = DB.AutoMigrate(&User{}, &Question{}, &Lang{}, &Answer{})
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	fmt.Println("Database connected and migrated")
}
