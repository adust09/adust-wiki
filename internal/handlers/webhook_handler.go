// internal/handlers/webhook_handler.go
package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"go-todo/internal/github"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request body", http.StatusInternalServerError)
		return
	}

	// Webhookのペイロードをコンソールに表示
	fmt.Printf("Received Webhook: %s\n", string(body))

	// GitHub APIからファイル取得のロジックを呼び出す（例）
	content, err := github.GetFileFromGitHub("username/repo", "path/to/file.md")
	if err != nil {
		http.Error(w, "Failed to get file from GitHub", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Received file content: %s", content)
}
