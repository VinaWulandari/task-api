package main

import (
	"log"
	"task-api/internal/database"
	"task-api/internal/models"
	"task-api/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "task-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task Management API
// @version 1.0
// @description Simple Task Management API with JWT Authentication
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment")
	}

	r := gin.Default()

	if err := database.ConnectDB(); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Auto migrate model
	database.DB.AutoMigrate(&models.User{}, &models.Task{})

	// Root test
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "🚀 Task API is running!"})
	})

	// Register auth routes
	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
