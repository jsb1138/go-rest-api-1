package database

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	utils "github.com/jsb1138/go-rest-2/utils"
)

const (
	host     string = "aurora-database-6-instance-1.cdmtvxl4golh.eu-central-1.rds.amazonaws.com"
	port     int    = 5432
	user     string = "postgres"
	password string = "passgres123"
	dbname   string = "postgres"
)

func DB(c *gin.Context) *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	utils.CheckError(err)

	// defer db.Close()
	return db
}
