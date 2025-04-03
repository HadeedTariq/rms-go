package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetDatabaseURL returns the database connection string
func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}
