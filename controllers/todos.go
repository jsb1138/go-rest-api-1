package controllers

import (
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	db "github.com/jsb1138/go-rest-2/database"
	utils "github.com/jsb1138/go-rest-2/utils"
)

func Todos(c *gin.Context) {
	rows, e := db.DB(c).Query("SELECT * FROM todos")
	utils.CheckError(e)

	var id string
	var name string
	var description string

	for rows.Next() {
		e = rows.Scan(&id, &name, &description)
		utils.CheckError(e)
		c.JSON(200, gin.H{
			"id":          id,
			"name":        name,
			"description": description,
		})
	}
}
