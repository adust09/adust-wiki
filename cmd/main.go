package main

import (
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
		authRoutes.POST("/register", register)
		authRoutes.POST("/login", login)
	}

	imageRoutes := r.Group("/api")
	{
		imageRoutes.POST("/upload", uploadImage)
		imageRoutes.GET("/images", listImages)
		imageRoutes.GET("/images/:imageId", downloadImage)
	}

	r.Run(":8080")
}
