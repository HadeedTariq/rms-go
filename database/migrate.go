package database

import (
	"log"
	"rms-platform/models"
)

func MigrateTables() {
	if DB == nil {
		log.Fatal("Database connection is not initialized")
	}

	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database tables:", err)
	}

	log.Println("Database tables migrated successfully")
}
