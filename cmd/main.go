package main

import (
	"todo-list-golang/internal/domain/service"
	"todo-list-golang/internal/handler"
	"todo-list-golang/internal/infrastructure/repository"

	_ "todo-list-golang/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	repo := repository.NewInMemoryTodoRepo()

	svc := service.NewTodoService(repo)

	h := handler.NewTodoHandler(svc)

	r := gin.Default()

	v1 := r.Group("/api/v1/todos")
	{
		v1.POST("/", h.Create)
		v1.GET("/", h.GetAll)
		v1.GET("/:id", h.GetOne)
		v1.PUT("/:id", h.Update)
		v1.DELETE("/:id", h.Delete)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}