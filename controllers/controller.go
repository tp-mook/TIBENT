package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "イベント一覧を取得しました。",
	})
}
