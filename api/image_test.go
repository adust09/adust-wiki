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

func TestListImages(t *testing.T) {
	router := gin.Default()
	router.GET("/api/images", ListImages)

	// HTTPリクエストを作成
	req, _ := http.NewRequest("GET", "/api/images", nil)
	w := httptest.NewRecorder()

	// テストリクエストを送信
	router.ServeHTTP(w, req)

	// ステータスコード200を確認
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスにダミーの画像URLが含まれていることを確認
	assert.Contains(t, w.Body.String(), "go-todo-images.s3.amazonaws.com/image1.png")
}

func TestDownloadImage(t *testing.T) {
	router := gin.Default()
	router.GET("/api/images/:imageId", DownloadImage)

	// HTTPリクエストを作成（imageIdとして "testimage.png" を指定）
	req, _ := http.NewRequest("GET", "/api/images/testimage.png", nil)
	w := httptest.NewRecorder()

	// テストリクエストを送信
	router.ServeHTTP(w, req)

	// ステータスコード200を確認
	assert.Equal(t, http.StatusOK, w.Code)

	// 正しいS3 URLが生成されていることを確認
	assert.Contains(t, w.Body.String(), "https://go-todo-images.s3.amazonaws.com/testimage.png")
}
