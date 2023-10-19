package test

import (
	"github.com/jesper.norgard/Todo_go/model"
)

type MockTodoRepository struct {
	tasks []model.Task
}

func NewMockTodoRepository() *MockTodoRepository {
	return &MockTodoRepository{
		tasks: []model.Task{
			{Name: "One", Completed: false},
			{Name: "Two", Completed: false},
		},
	}
}

func (taskRepository *MockTodoRepository) InsertTask(task string) error {
	taskRepository.tasks = append(taskRepository.tasks, model.Task{Name: task, Completed: false})
	return nil
}

func (taskRepository *MockTodoRepository) GetTasks() ([]model.Task, error) {
	return taskRepository.tasks, nil
}

func (taskRepository *MockTodoRepository) CompleteTask(int) error {
	return nil
}
