package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ユーザー登録の構造体
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// ユーザー登録処理
func register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// パスワードをハッシュ化し、ユーザー情報をデータベースに保存する処理を実装
	// 仮にユーザーが作成されたとする
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"userId":  "sample-user-id", // 実際には生成されたユーザーIDを返す
	})
}
