package db

import (
	"imagera/internal/db/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := DB.AutoMigrate(&models.User{}, &models.Image{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")
}
