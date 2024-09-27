package api

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUploadImage(t *testing.T) {
	router := gin.Default()
	router.POST("/api/upload", UploadImage)

	fileContent := []byte("This is a test image")
	reqBody := new(bytes.Buffer)
	writer := multipart.NewWriter(reqBody)
	part, _ := writer.CreateFormFile("file", "testimage.png")
	part.Write(fileContent)
	writer.Close()

	req, _ := http.NewRequest("POST", "/api/upload", reqBody)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Image uploaded successfully")
}
