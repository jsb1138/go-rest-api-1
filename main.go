package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jsb1138/go-rest-2/routes"

	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()    // Instantiate the router
	router.Use(cors.Default()) // Enable CORS
	routes.Routes(router)      // Register the routes
	router.Run(":8080")        // Start the server
}
