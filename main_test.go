package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jesper.norgard/Todo_go/database"
	"github.com/jesper.norgard/Todo_go/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasksRoute(t *testing.T) {
	todoRepository := test.NewMockTodoRepository()
	router := SetupRouter(todoRepository)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	router.ServeHTTP(w, req)

	// todo add expected testing

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestToggleCompleteByIdRoute(t *testing.T) {
	todoRepository := test.NewMockTodoRepository()
	router := SetupRouter(todoRepository)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTasksFromDb(t *testing.T) {
	Db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer Db.Close()

	mock.ExpectQuery("SELECT \\* FROM tasks").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "completed"}).
			AddRow(1, "Task 1", true))

	tasks, err := database.GetTasks(Db)
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	assert.Len(t, tasks, 1)
	assert.Equal(t, 1, tasks[0].ID)
	assert.Equal(t, "Task 1", tasks[0].Name)
	assert.True(t, tasks[0].Completed)
}
