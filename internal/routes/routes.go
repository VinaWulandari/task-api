package routes

import (
	"task-api/internal/handlers"
	"task-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)

		// ✅ Task routes (JWT Protected)
		taskRoutes := api.Group("/tasks")
		taskRoutes.Use(middleware.AuthMiddleware())
		{
			taskRoutes.POST("/", handlers.CreateTask)
			taskRoutes.GET("/", handlers.GetTasks)
			taskRoutes.GET("/:id", handlers.GetTask)
			taskRoutes.PUT("/:id", handlers.UpdateTask)
			taskRoutes.DELETE("/:id", handlers.DeleteTask)
		}

	}
}
