// cmd/webhook/main.go
package main

import (
	"fmt"
	"net/http"

	"go-todo/internal/handlers"
)

func main() {
	http.HandleFunc("/webhook", handlers.WebhookHandler) // Webhookリクエストの処理
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
