package handlers

import (
	"sync"

	"go-todo/api"
)

func AnalyzeADRsParallel(adrs []string, openAIKey string) []string {
	var wg sync.WaitGroup
	results := make([]string, len(adrs))

	for i, adr := range adrs {
		wg.Add(1)
		go func(i int, adr string) {
			defer wg.Done()
			result, _ := api.AnalyzeWithOpenAI(openAIKey, adr)
			results[i] = result
		}(i, adr)
	}

	wg.Wait()
	return results
}

func ScheduleADRAnalysis() {
	adrs, _ := api.FetchADRFromNotion("your-notion-api-key", "database-id")
	results := AnalyzeADRsParallel(adrs, "your-openai-api-key")

	for _, result := range results {
		// Save result to Notion
	}
}
