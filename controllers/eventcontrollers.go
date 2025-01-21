package controllers

import (
	"net/http"
	"strconv"

	"github.com/username/TIBENT/database"
	"github.com/username/TIBENT/models"

	"github.com/gin-gonic/gin"
)

// イベント一覧を取得
func GetEvents(c *gin.Context) {
	var events []models.Event
	database.DB.Find(&events)
	c.JSON(http.StatusOK, events)
}

// イベントを作成
func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&event)
	c.JSON(http.StatusOK, event)
}

// 特定のイベントを取得
func GetEventByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var event models.Event
	if result := database.DB.First(&event, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

// イベントを更新
func UpdateEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var event models.Event
	if result := database.DB.First(&event, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&event)
	c.JSON(http.StatusOK, event)
}

// イベントを削除
func DeleteEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var event models.Event
	if result := database.DB.First(&event, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	database.DB.Delete(&event)
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
