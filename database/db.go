package database

import (
	"fmt"
	"log"
	"rms-platform/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {

	config.LoadEnv()

	dsn := config.GetDatabaseURL()
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in the .env file")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Error getting raw DB instance: ", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	fmt.Println("Connected to the database!")
}
