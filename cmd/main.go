package main

import (
	"imagera/api"
	"imagera/internal/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// データベースに接続
	err := db.Connect() // これでグローバル変数DBに接続情報がセットされる
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// マイグレーションを実行
	db.Migrate()

	// Ginのデフォルトルーターを作成
	r := gin.Default()

	// APIのヘルスチェックエンドポイント
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "API is up and running!"})
	})

	// 認証系のルートグループ
	authRoutes := r.Group("/api")
	{
		// それぞれのエンドポイントでRegister, Login, Dashboardを呼び出し
		authRoutes.POST("/register", api.Register)
		authRoutes.POST("/login", api.Login)
		authRoutes.GET("/dashboard", api.Dashboard)
	}

	// サーバーを8080ポートで起動
	r.Run(":8080")
}
