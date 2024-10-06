package main

import (
	"imagera/api"
	"imagera/internal/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func (app *App) Register(c *gin.Context) {
	api.Register(c)
}

func (app *App) Login(c *gin.Context) {
	api.Login(c)
}

func (app *App) Dashboard(c *gin.Context) {
	api.Dashboard(c)
}

func main() {
	database := &db.GormDB{}
	conn, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	app := &App{DB: conn}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "API is up and running!"})
	})

	authRoutes := r.Group("/api")
	{
		authRoutes.POST("/register", app.Register)
		authRoutes.POST("/login", app.Login)
		authRoutes.GET("/dashboard", app.Dashboard)
	}

	r.Run(":8080")
}
