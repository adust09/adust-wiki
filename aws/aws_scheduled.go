package aws

import (
	"go-todo/api"
	"go-todo/handlers"
)

func ScheduleADRAnalysis() {
	adrs, _ := api.FetchADRFromNotion("your-notion-api-key", "database-id")
	results := handlers.AnalyzeADRsParallel(adrs, "your-openai-api-key")

	for _, result := range results {
		// Save result to Notion
	}
}
