package routes

import (
	//"TIBENT/controllers"
	"github.com/username/TIBENT/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/events", controllers.GetEvents)
		api.POST("/events", controllers.CreateEvent)
		api.GET("/events/:id", controllers.GetEventByID)
		api.PUT("/events/:id", controllers.UpdateEvent)
		api.DELETE("/events/:id", controllers.DeleteEvent)
	}
}
