package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	db "github.com/jsb1138/go-rest-2/database"
)

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created_At  string `json:"created_at"`
	Updated_At  string `json:"updated_at"`
}

func Todos(c *gin.Context) {
	// Build the query with squirrel
	query, args, err := squirrel.Select("*").From("todos").OrderBy("created_at ASC").ToSql()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Execute the query on the database
	rows, err := db.DB().Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Create a slice of todos
	var todos []Todo

	// Iterate over the rows, adding the individual todos to the slice
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Created_At, &todo.Updated_At)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert the slice of todos to JSON
	jsonData, err := json.Marshal(todos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Set the content type and status code, then write the JSON data to the response
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(jsonData)
}

func CreateTodo(c *gin.Context) {
	// Declare a Todo object
	var todo Todo
	// Parse the request body into a Todo object
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Execute the INSERT query on the database
	_, err = db.DB().Exec("INSERT INTO todos VALUES ($1, $2, $3, $4, $5)", todo.ID, todo.Title, todo.Description, todo.Created_At, todo.Updated_At)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the status and the created todo as a JSON response for confirmation
	c.JSON(http.StatusOK, todo)
}

func EditTodo(c *gin.Context) {
	// Get the todo ID from the URL parameter
	id := c.Param("id")

	// Declare a Todo object
	var todo Todo
	// Parse the request body into a Todo object
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Execute the UPDATE query on the database
	_, err = db.DB().Exec("UPDATE todos SET title=$1, description=$2, updated_at=$3 WHERE id=$4", todo.Title, todo.Description, todo.Updated_At, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the status and the updated todo as a JSON response for confirmation
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	// Get the todo ID from the URL parameter
	id := c.Param("id")

	// Execute the DELETE query on the database
	_, err := db.DB().Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the status and a message as a JSON response for confirmation
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func DeleteTodos(c *gin.Context) {
	// Declare an empty slice of todo IDs
	var ids []string

	// Get the todo IDs as a string from the URL parameter
	params := c.Param("ids")
	// Split the string into a slice of IDs
	ids = strings.Split(params, ",")

	// Structure the query string by iterating over the slice of IDs
	query := "DELETE FROM todos WHERE id IN ("
	for i, id := range ids {
		if i > 0 {
			query += ", "
		}
		query += "'" + id + "'"
	}
	query += ")"

	// Execute the DELETE query on the database
	res, err := db.DB().Exec(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Confirm the number of todos deleted
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the status and a message as a JSON response for confirmation of the number of todos deleted
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d todos deleted successfully", rowsAffected),
	})
}
