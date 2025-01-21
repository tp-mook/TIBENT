package main

import (
	"github.com/username/TIBENT/database"
	"github.com/username/TIBENT/models"
	"github.com/username/TIBENT/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// データベース接続
	database.ConnectDB()

	// データベースのマイグレーション
	database.DB.AutoMigrate(&models.Event{})

	// ルートの設定
	r := gin.Default()
	routes.SetupRoutes(r)

	// サーバー起動
	r.Run(":8080")
}
