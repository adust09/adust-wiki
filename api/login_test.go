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

// テスト用のルーターセットアップ
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	r.POST("/logout", api.Logout)
	r.GET("/dashboard", api.AuthMiddleware(), api.Dashboard)
	return r
}

// ユーザー登録のテスト
func TestRegister(t *testing.T) {
	r := setupRouter()

	// リクエストデータ作成
	userData := map[string]string{
		"email":    "testuser@example.com",
		"password": "testpassword",
	}
	jsonData, _ := json.Marshal(userData)

	// POSTリクエスト送信
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ステータスコードとレスポンスの検証
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Registration successful")
}

// ログインテスト
func TestLogin(t *testing.T) {
	r := setupRouter()

	// ユーザー登録
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

	// ステータスコードとレスポンスの検証
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Login successful")
}

// ログアウトのテスト
func TestLogout(t *testing.T) {
	r := setupRouter()

	// ユーザー登録とログイン
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

	// ログアウトリクエスト
	req, _ = http.NewRequest("POST", "/logout", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ステータスコードとレスポンスの検証
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Logout successful")
}

// 認証なしで保護されたページにアクセスした場合のテスト
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

// 認証済みで保護されたページにアクセスした場合のテスト
func TestProtectedPageWithLogin(t *testing.T) {
	r := setupRouter()

	// ユーザー登録とログイン
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

	// 認証済み状態でダッシュボードにアクセス
	req, _ = http.NewRequest("GET", "/dashboard", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ステータスコードとメッセージの検証
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Welcome to your dashboard")
}
