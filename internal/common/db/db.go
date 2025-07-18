package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	database *gorm.DB
	once     sync.Once
)

func ConnectDB() *gorm.DB {
	once.Do(func() {
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || dbName == "" {
			log.Fatal("❌ One or more required database environment variables are not set")
		}

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			dbHost, dbUser, dbPass, dbName, dbPort,
		)

		var err error
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Fatalf("❌ Failed to connect to the database: %v", err)
		}

		log.Println("📊 Database has been successfully connected")
	})

	return database
}
