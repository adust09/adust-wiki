// webhook.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request body", http.StatusInternalServerError)
		return
	}

	// Webhookのペイロードを処理するロジックをここに追加
	fmt.Printf("Received webhook: %s\n", string(body))

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}
