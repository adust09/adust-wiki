package main

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

func main() {
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Article{})

	article := Article{Title: "サンプル記事", Content: "この記事の内容", Tags: []string{"Go", "AWS"}, Date: "2024-09-20"}
	db.Create(&article)

}
