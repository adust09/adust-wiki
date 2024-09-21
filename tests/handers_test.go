package handlers_test

import (
	"go-todo/internal/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// 正常なWebhookリクエストのテスト
func TestWebhookHandler_Success(t *testing.T) {
	// Webhookリクエストのボディを模擬
	requestBody := `{
        "ref": "refs/heads/main",
        "repository": {
            "name": "sample-repo",
            "owner": {
                "name": "testuser"
            }
        },
        "head_commit": {
            "id": "abc123",
            "message": "Test commit"
        }
    }`

	req, err := http.NewRequest("POST", "/webhook", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// レコーダーでレスポンスをキャプチャ
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.WebhookHandler)

	handler.ServeHTTP(rr, req)

	// ステータスコードが200（OK）か確認
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	// レスポンスボディの確認（必要に応じて）
	expectedResponse := "Received Webhook"
	if rr.Body.String() != expectedResponse {
		t.Errorf("unexpected body: got %v, want %v", rr.Body.String(), expectedResponse)
	}
}

// 不正なメソッドに対するテスト
func TestWebhookHandler_InvalidMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/webhook", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.WebhookHandler)

	handler.ServeHTTP(rr, req)

	// ステータスコードが405（Method Not Allowed）か確認
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("expected status code %d, got %d", http.StatusMethodNotAllowed, status)
	}
}
