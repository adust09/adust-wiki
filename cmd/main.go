package main

import (
	"go-todo/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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
	}

	imageRoutes := r.Group("/api")
	{
		imageRoutes.POST("/upload", api.UploadImage)
	}

	r.Run(":8080")
}
