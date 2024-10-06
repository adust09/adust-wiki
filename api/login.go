package api

import (
	"imagera/internal/db"
	"imagera/internal/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

var req struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
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

func Login(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !checkPasswordHash(req.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	session.Values["authenticated"] = true
	session.Values["username"] = user.Email
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Logout(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")

	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := store.Get(c.Request, "session-name")

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func Dashboard(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")
	username := session.Values["username"].(string)
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to your dashboard, " + username})
}
