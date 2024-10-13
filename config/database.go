package config

import (
	"fmt"
	"log"

	"github.com/R-Thibault/OrgaJobSearch/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := GetConfig("DB_HOST")
	user := GetConfig("DB_USER")
	password := GetConfig("DB_PASSWORD")
	dbName := GetConfig("DB_NAME")
	port := GetConfig("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Connected to the database successfully!")

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}
	log.Println("Database schema migrated successfully!")
}

func CloseDB() {
	if sqlDB, err := DB.DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing the database connection: %v", err)
		} else {
			log.Println("Database connection closed.")
		}
	} else {
		log.Printf("Error retrieving the generic database object: %v", err)
	}
}
