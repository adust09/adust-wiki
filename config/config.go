// config/config.go
package config

import (
	"fmt"
	"os"
)

func LoadEnv() {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		fmt.Println("Warning: GITHUB_TOKEN is not set")
	}
}
