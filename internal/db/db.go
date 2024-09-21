// internal/db/db.go
package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Article struct {
	ID      uint `gorm:"primaryKey"`
	Title   string
	Content string
	Tags    []string `gorm:"type:text[]"`
	Date    string
}

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Article{})
	return db, nil
}

func SaveArticle(db *gorm.DB, article *Article) error {
	return db.Create(article).Error
}
