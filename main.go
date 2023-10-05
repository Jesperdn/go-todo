package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jesper.norgard/Todo_go/database"
	"github.com/jesper.norgard/Todo_go/model"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

func main() {
	Db := database.SetupDb()
	router := setupRouter(Db)
	router.Run("localhost:8080")
}

func setupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/tasks", func(c *gin.Context) {
		getTasksHandler(c, db)
	})
	r.POST("/tasks", func(c *gin.Context) {
		postTasksHandler(c, db)
	})
	r.PUT("/tasks", func(c *gin.Context) {
		toggleCompleteByIdHandler(c, db)
	})

	return r
}

func getTasksHandler(c *gin.Context, db *sql.DB) {
	tasks, err := database.GetTasks(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func postTasksHandler(c *gin.Context, db *sql.DB) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.PostTask(db, &newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func toggleCompleteByIdHandler(c *gin.Context, db *sql.DB) {
	idParam := c.Query("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param id not given"})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err = database.ToggleCompleteById(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}
