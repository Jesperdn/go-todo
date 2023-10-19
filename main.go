package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jesper.norgard/Todo_go/biz/todo"
	"github.com/jesper.norgard/Todo_go/database"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"

	todoController "github.com/jesper.norgard/Todo_go/rest/todo"
)

func main() {
	Db := database.SetupDb()
	defer Db.Close()

	todoRepository := todo.NewTaskRepository(Db)

	router := SetupRouter(todoRepository)
	router.Run()
}

func SetupRouter(taskRepository todo.TaskRepositoryContract) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/tasks", func(c *gin.Context) {
		todoController.Get(c, taskRepository)
	})
	r.POST("/tasks", func(c *gin.Context) {
		todoController.Post(c, taskRepository)
	})
	r.PUT("/tasks", func(c *gin.Context) {
		todoController.Update(c, taskRepository)
	})

	return r
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
