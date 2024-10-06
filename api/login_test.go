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

// テスト用のGinエンジンをセットアップ
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	r.POST("/logout", api.Logout)
	r.GET("/dashboard", api.AuthMiddleware(), api.Dashboard)
	return r
}

func TestRegister(t *testing.T) {
	r := setupRouter()

	// リクエストデータの作成
	userData := map[string]string{
		"email":    "testuser@example.com",
		"password": "testpassword",
	}
	jsonData, _ := json.Marshal(userData)

	// POSTリクエストを送信
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ステータスコードとレスポンスメッセージの検証
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Registration successful")
}

func TestLogin(t *testing.T) {
	r := setupRouter()

	// まずユーザーを登録する
	userData := map[string]string{
		"email":    "testuser@example.com",
		"password": "testpassword",
	}
	jsonData, _ := json.Marshal(userData)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ログインリクエスト
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ステータスコードとレスポンスメッセージの検証
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Login successful")
}

func TestLogout(t *testing.T) {
	r := setupRouter()

	// まずユーザーを登録してログインする
	userData := map[string]string{
		"email":    "testuser@example.com",
		"password": "testpassword",
	}
	jsonData, _ := json.Marshal(userData)

	// ユーザー登録
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ログイン
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ログアウト
	req, _ = http.NewRequest("POST", "/logout", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ステータスコードとレスポンスメッセージの検証
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Logout successful")
}

func TestProtectedPageWithoutLogin(t *testing.T) {
	r := setupRouter()

	// ログインせずにダッシュボードにアクセス
	req, _ := http.NewRequest("GET", "/dashboard", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ステータスコードとエラーメッセージの検証
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Unauthorized")
}

func TestProtectedPageWithLogin(t *testing.T) {
	r := setupRouter()

	// まずユーザーを登録してログインする
	userData := map[string]string{
		"email":    "testuser@example.com",
		"password": "testpassword",
	}
	jsonData, _ := json.Marshal(userData)

	// ユーザー登録
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ログイン
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 認証済みでダッシュボードにアクセス
	req, _ = http.NewRequest("GET", "/dashboard", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ステータスコードとメッセージの検証
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Welcome to your dashboard")
}
