package todo

import (
	"database/sql"
	"github.com/jesper.norgard/Todo_go/model"
)

type TaskRepositoryContract interface {
	InsertTask(string) error
	GetTasks() ([]model.Task, error)
	CompleteTask(int) error
}

type TaskRepository struct {
	db *sql.DB
}

func (taskRepository *TaskRepository) InsertTask(task string) error {
	_, err := taskRepository.db.Exec("INSERT INTO tasks (name, completed) VALUES ($1, $2)", task, false)
	return err
}

func (taskRepository *TaskRepository) GetTasks() ([]model.Task, error) {
	rows, err := taskRepository.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (taskRepository *TaskRepository) CompleteTask(id int) error {
	_, err := taskRepository.db.Exec("UPDATE tasks SET completed = NOT completed WHERE id = $1", id)
	return err
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}
