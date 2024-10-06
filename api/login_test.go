package api_test

import (
	"bytes"
	"encoding/json"
	"imagera/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	r.POST("/logout", api.Logout)
	r.GET("/dashboard", api.AuthMiddleware(), api.Dashboard)
	return r
}

func TestLogin(t *testing.T) {
	r := setupRouter()

	userData := map[string]string{
		"email":    "testuser@example.com",
		"password": "testpassword",
	}
	jsonData, _ := json.Marshal(userData)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Login successful")
}

func TestLogout(t *testing.T) {
	r := setupRouter()

	userData := map[string]string{
		"email":    "testuser@example.com",
		"password": "testpassword",
	}
	jsonData, _ := json.Marshal(userData)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("POST", "/logout", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Logout successful")
}

func TestProtectedPageWithoutLogin(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/dashboard", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Unauthorized")
}

func TestProtectedPageWithLogin(t *testing.T) {
	r := setupRouter()

	userData := map[string]string{
		"email":    "testuser@example.com",
		"password": "testpassword",
	}
	jsonData, _ := json.Marshal(userData)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/dashboard", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Welcome to your dashboard")
}
