package routes

import (
	"goGin/handlers"
	"goGin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.RegisterUser)
		auth.POST("/login", handlers.LoginUser)
	}
	
	tasks := api.Group("/tasks")
	tasks.Use(middleware.RequiredAuth)
	{
		tasks.GET("", handlers.GetTasks)
		tasks.POST("", handlers.CreateTask)
		tasks.GET("/:id", handlers.GetSingleTask)
		tasks.PUT("/:id", handlers.UpdateTask)
		tasks.PATCH("/:id", handlers.PatchTask)
		tasks.DELETE("/:id", handlers.DeleteTask)
	}
}