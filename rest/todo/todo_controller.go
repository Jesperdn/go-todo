package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/jesper.norgard/Todo_go/biz/todo"
	"github.com/jesper.norgard/Todo_go/model"
	"net/http"
	"strconv"
)

func Get(c *gin.Context, todoRepository todo.TaskRepositoryContract) {
	tasks, err := todoRepository.GetTasks()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, tasks)
}

func Post(c *gin.Context, todoRepository todo.TaskRepositoryContract) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todoRepository.InsertTask(newTask.Name)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func Update(c *gin.Context, todoRepository todo.TaskRepositoryContract) {
	idParam := c.Query("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param id not given"})
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = todoRepository.CompleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}
