package routes

import (
	"goGin/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		tasks := api.Group("/tasks")
		{
			tasks.GET("", handlers.GetTasks)
			tasks.POST("", handlers.CreateTask)
			tasks.GET("/:id", handlers.GetTask)
		}
	}
}
