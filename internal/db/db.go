package db

import (
	"fmt"
	"imagera/internal/db/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// グローバル変数としてgorm.DBを保持
var DB *gorm.DB

// Connect - PostgreSQLデータベースに接続
func Connect() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	log.Println("Successfully connected to PostgreSQL database")
	return nil
}

// Migrate - テーブルのマイグレーションを実行
func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Image{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")
}
