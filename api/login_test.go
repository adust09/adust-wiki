package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	router := gin.Default()
	router.POST("/api/login", Login)

	requestBody := `{
		"email": "test@example.com",
		"password": "password123"
	}`

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Contains(t, w.Body.String(), "Login successful")
	assert.Contains(t, w.Body.String(), "token")
}

func TetLoginInvalidEmail(t *testing.T) {
	router := gin.Default()
	router.POST("/api/login", Login)

	requestBody := `{
		"email": "invalid-email",
		"password": "password123"
	}`

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestLoginMissingPassword(t *testing.T) {
	router := gin.Default()
	router.POST("/api/login", Login)

	requestBody := `{
		"email": "test@example.com"
	}`

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	assert.Contains(t, w.Body.String(), "error")
}
