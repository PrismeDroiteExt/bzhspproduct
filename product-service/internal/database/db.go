package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"bzhspback.fr/breizhsport/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	var err error
	maxRetries := 5
	retryInterval := time.Second * 5

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("PRODUCT_DB_HOST"),
		os.Getenv("PRODUCT_DB_USER"),
		os.Getenv("PRODUCT_DB_PASSWORD"),
		os.Getenv("PRODUCT_DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Printf("Successfully connected to database on attempt %d", i+1)
			break
		}
		log.Printf("Failed to connect to database, attempt %d/%d: %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(retryInterval)
		}
	}

	if err != nil {
		return fmt.Errorf("failed to connect to product database after %d attempts: %v", maxRetries, err)
	}
	// Auto migrate models
	return db.AutoMigrate(&models.Product{}, &models.Brand{}, &models.Category{})
}

func GetDB() *gorm.DB {
	return db
}
