package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/username/TIBENT/auth"
	"github.com/username/TIBENT/controllers"
)

// APIルート設定
func SetupRoutes(r *gin.Engine) {
	// 認証が不要なルート
	r.POST("/api/login", auth.Login)

	// 認証が必要なルート (JWT認証ミドルウェアを適用)
	authorized := r.Group("/api/events")
	authorized.Use(auth.AuthMiddleware())
	{
		authorized.GET("", controllers.GetEvents)          // イベント一覧取得
		authorized.GET("/:id", controllers.GetEventByID)   // 特定のイベント取得
		authorized.POST("", controllers.CreateEvent)       // イベント作成
		authorized.PUT("/:id", controllers.UpdateEvent)    // イベント更新
		authorized.DELETE("/:id", controllers.DeleteEvent) // イベント削除
	}
}
