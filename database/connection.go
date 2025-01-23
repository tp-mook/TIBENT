package database

import (
	"fmt"
	"log"
	"os"

	"github.com/username/TIBENT/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB: グローバルなデータベースインスタンス
var DB *gorm.DB

// ConnectDB: データベースに接続する関数
func ConnectDB() {
	// 環境変数から接続情報を取得
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// データベース接続
	var dbErr error
	DB, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatalf("Failed to connect to database: %v", dbErr)
	}

	log.Println("Database connection established")

	// マイグレーションの実行
	err := DB.AutoMigrate(&models.Event{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Database migration completed")

}
