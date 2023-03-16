package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/jsb1138/go-rest-2/controllers"
)

func Routes(router *gin.Engine) {

	router.GET("/todos", controllers.Todos)
	// router.POST("/todo", controllers.CreateTodo)
	// router.GET("/todo/:todoId", controllers.GetSingleTodo)
	// router.PUT("/todo/:todoId", controllers.EditTodo)
	// router.DELETE("/todo/:todoId", controllers.DeleteTodo)

	router.GET("/", Connected)
	router.NoRoute(notFound)
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
}

func Connected(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Successful server connection",
	})
}
