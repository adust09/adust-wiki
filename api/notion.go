package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Notion APIを使ってADRを取得
func FetchADRFromNotion(apiKey string, databaseID string) (string, error) {

	notionAPIURL := os.Getenv("NOTION_API_URL")
	if notionAPIURL == "" {
		return "", fmt.Errorf("Notion API URL not set")
	}

	reqBody := []byte(`{}`)
	req, err := http.NewRequest("POST", notionAPIURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2021-08-16")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
