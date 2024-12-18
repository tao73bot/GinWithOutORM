package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-postgres/controllers"
)

func SetupTodoRoutes(router *gin.Engine, todoController *controllers.TodoController) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/todos", todoController.CreateTodo)
		v1.GET("/todos", todoController.GetTodos)
		v1.GET("/todos/:id", todoController.GetTodo)
		v1.PUT("/todos/:id", todoController.UpdateTodo)
		v1.DELETE("/todos/:id", todoController.DeleteTodo)
	}
}
