package main

import (
	"imagera/api"
	"imagera/internal/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	database := &db.GormDB{}

	_, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	db.Migrate()

	log.Println("Application started successfully")
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "API is up and running!",
		})
	})

	authRoutes := r.Group("/api")
	{
		authRoutes.POST("/register", api.Register)
		authRoutes.POST("/login", api.Login)
		authRoutes.POST("/logout", api.Logout)
		authRoutes.GET("/dashboard", api.Dashboard)

	}

	imageRoutes := r.Group("/api")
	{
		imageRoutes.POST("/upload", api.UploadImage)
	}

	r.Run(":8080")
}
