package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jesper.norgard/Todo_go/biz/todo"
	"github.com/jesper.norgard/Todo_go/database"
	"github.com/jesper.norgard/Todo_go/rest/middleware"
	todoController "github.com/jesper.norgard/Todo_go/rest/todo"
	_ "github.com/lib/pq"
	"net/http"
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
	r.Use(middleware.CORSMiddleware())

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
