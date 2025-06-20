package routes

import (
	"todo-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos", controllers.UpdateTodo)
	r.Delete("/todos", controllers.DeleteTodo)
}
