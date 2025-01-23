package main

import (
	"log"

	"github.com/username/TIBENT/database"
	"github.com/username/TIBENT/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .env ファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// データベース接続
	database.ConnectDB()

	// サーバー設定
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
