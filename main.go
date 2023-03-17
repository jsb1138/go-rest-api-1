package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jsb1138/go-rest-2/routes"

	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "PUT", "DELETE", "POST"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	routes.Routes(router)
	router.Run(":8080")
}
