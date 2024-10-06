package api

import (
	"imagera/internal/db"
	"imagera/internal/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var req struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser := models.User{
		Email:        req.Email,
		PasswordHash: string(passwordHash),
	}

	if err := db.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	session, _ := store.Get(c.Request, "session-name")
	session.Values["authenticated"] = true
	session.Values["username"] = req.Email
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
