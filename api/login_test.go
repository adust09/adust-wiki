package api_test

import (
	"imagera/api"
	"imagera/internal/db"
	"imagera/internal/db/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setup() {
	db.Connect()
	db.DB.Exec("DELETE FROM users")
}

func createTestContext(method, path string, body string) (*gin.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	return c, w
}

func TestRegister(t *testing.T) {
	setup()

	body := `{"username": "testuser", "email": "testuser@example.com", "password": "password123"}`
	c, w := createTestContext("POST", "/register", body)

	api.Register(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User registered successfully")
}

func TestLogin(t *testing.T) {
	setup()

	hashedPassword, _ := api.HashPassword("password123")
	user := models.User{
		Username:     "testuser",
		Email:        "testuser@example.com",
		PasswordHash: hashedPassword,
	}
	db.DB.Create(&user)

	body := `{"email": "testuser@example.com", "password": "password123"}`
	c, w := createTestContext("POST", "/login", body)

	api.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Login successful")
	assert.Contains(t, w.Body.String(), "token")
}
