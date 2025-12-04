package main

import (
	"log"

	"github.com/Amirali-nourian/Todo-list-golang/internal/config"
	"github.com/Amirali-nourian/Todo-list-golang/internal/domain"
	"github.com/Amirali-nourian/Todo-list-golang/internal/domain/repository"
	"github.com/Amirali-nourian/Todo-list-golang/internal/domain/service"
	"github.com/Amirali-nourian/Todo-list-golang/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Amirali-nourian/Todo-list-golang/docs"
)

// @title Todo List API
// @version 1.0
// @description A simple Todo List API with Swagger documentation
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Initialize database
	db := config.InitDB()

	// Auto migrate the schema
	if err := db.AutoMigrate(&domain.Todo{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize layers
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	// Setup Gin router
	router := gin.Default()

	// سرو کردن فایل index.html در صفحه اصلی
	router.StaticFile("/", "./web/index.html")
	// -------------------------

	// Swagger endpoint
	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	v1 := router.Group("/api/v1")
	{
		todos := v1.Group("/todos")
		{
			todos.POST("", todoHandler.CreateTodo)
			todos.GET("", todoHandler.GetAllTodos)
			todos.GET("/:id", todoHandler.GetTodoByID)
			todos.PUT("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
		}
	}

	// Start server
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
