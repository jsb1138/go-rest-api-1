package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/jsb1138/go-rest-2/controllers"
)

func Routes(router *gin.Engine) {
	// Define the routes
	router.GET("/todos", controllers.Todos)
	router.POST("/todo", controllers.CreateTodo)
	router.PUT("/todo/:id", controllers.EditTodo)
	router.DELETE("/todo/:id", controllers.DeleteTodo)
	router.DELETE("/todos/:ids", controllers.DeleteTodos)

	// Handle extraneous routes
	router.GET("/", Connected)
	router.NoRoute(notFound)
}

// 404 route handler
func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
}

// Bare endpoint to test server connection
func Connected(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Successful server connection",
	})
}
