package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	db "github.com/jsb1138/go-rest-2/database"
	utils "github.com/jsb1138/go-rest-2/utils"
)

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ID struct {
	ID string `json:"ids"`
}

func Todos(ctx *gin.Context) {
	rows, err := db.DB().Query("SELECT * FROM todos")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description)
		utils.CheckError(err)
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jsonData, err := json.Marshal(todos)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write(jsonData)
}

func CreateTodo(ctx *gin.Context) {
	var todo Todo
	err := ctx.BindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.DB().Exec("INSERT INTO todos VALUES ($1, $2, $3)", todo.ID, todo.Title, todo.Description)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func EditTodo(ctx *gin.Context) {
	// Get the todo ID from the URL parameter
	id := ctx.Param("id")

	// Parse the request body into a Todo object
	var todo Todo
	err := ctx.BindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Execute the UPDATE query on the database
	_, err = db.DB().Exec("UPDATE todos SET title=$1, description=$2 WHERE id=$3", todo.Title, todo.Description, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the updated todo as a JSON response
	ctx.JSON(http.StatusOK, todo)
}

func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := db.DB().Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func DeleteTodos(ctx *gin.Context) {
	var ids []string

	params := ctx.Param("ids")
	ids = strings.Split(params, ",")

	query := "DELETE FROM todos WHERE id IN ("
	for i, id := range ids {
		if i > 0 {
			query += ", "
		}
		query += "'" + id + "'"
	}
	query += ")"

	res, err := db.DB().Exec(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d todos deleted successfully", rowsAffected),
	})
}
