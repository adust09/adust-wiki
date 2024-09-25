package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのデフォルトのルーターを作成
	r := gin.Default()

	// ヘルスチェック用エンドポイント
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "API is up and running!",
		})
	})

	// 認証関連エンドポイントのグループ
	authRoutes := r.Group("/api")
	{
		authRoutes.POST("/register", register)
		authRoutes.POST("/login", login)
	}

	// 画像関連エンドポイントのグループ
	imageRoutes := r.Group("/api")
	{
		imageRoutes.POST("/upload", uploadImage)
		imageRoutes.GET("/images", listImages)
		imageRoutes.GET("/images/:imageId", downloadImage)
	}

	// サーバーを起動
	r.Run(":8080") // ポート8080でリッスン
}
