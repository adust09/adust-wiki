package aws

import (
	"project-root/api"
	"project-root/handlers"
)

// AWSのスケジュールタスク処理
func ScheduleADRAnalysis() {
	// Notionからデータを取得し、解析して結果を保存
	adrs, _ := api.FetchADRFromNotion("your-notion-api-key", "database-id")
	results := handlers.AnalyzeADRsParallel(adrs, "your-openai-api-key")

	// 結果を別のNotionページに保存（Notion APIで出力）
	for _, result := range results {
		// Save result to Notion
	}
}
