package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Start of Selection
// OpenAI APIを使ってADRを解析
func AnalyzeWithOpenAI(apiKey, adrText string) (string, error) {
	openAIURL := os.Getenv("OPENAI_API_URL")
	if openAIURL == "" {
		return "", fmt.Errorf("OpenAI API URL not set")
	}

	reqBody := map[string]interface{}{
		"prompt":    adrText,
		"max_token": 100,
	}

	jsonBody, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", openAIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

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

	return string(body), nil // Placeholder return to fix missing return error
}
