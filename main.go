package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jsb1138/go-rest-2/routes"

	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	routes.Routes(router)

	router.Run(":8080")

	// insert
	// insertStmt := `INSERT INTO todos VALUES ($1, $2, $3);`
	// _, e := db.Exec(insertStmt, "BbB888888881Hw3333333", "test", "testing!")
	// CheckError(e)
}
