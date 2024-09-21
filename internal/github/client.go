// internal/github/client.go
package github

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetFileFromGitHub(repo, path string) ([]byte, error) {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN is not set")
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repo, path)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+githubToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch file: %s", resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
