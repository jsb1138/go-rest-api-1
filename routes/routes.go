package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/jsb1138/go-rest-2/controllers"
)

func Routes(router *gin.Engine) {

	router.GET("/todos", controllers.Todos)

	router.GET("/", controllers.Welcome)
	// router.GET("/", welcome)
	// router.GET("/todos", controllers.GetAllTodos)
	// router.POST("/todo", controllers.CreateTodo)
	// router.GET("/todo/:todoId", controllers.GetSingleTodo)
	// router.PUT("/todo/:todoId", controllers.EditTodo)
	// router.DELETE("/todo/:todoId", controllers.DeleteTodo)
	router.NoRoute(notFound)
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
}
