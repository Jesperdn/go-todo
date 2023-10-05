package database

import (
	"database/sql"
	"fmt"
	"github.com/jesper.norgard/Todo_go/model"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "jesper.norgard"
	dbname = "go_test_db"
)

func SetupDb() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	return db
}

func GetTasks(db *sql.DB) ([]model.Task, error) {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func PostTask(db *sql.DB, task *model.Task) (sql.Result, error) {
	res, err := db.Exec("INSERT INTO tasks (name, completed) VALUES ($1, $2)", task.Name, task.Completed)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func ToggleCompleteById(db *sql.DB, id int) (sql.Result, error) {
	res, err := db.Exec("UPDATE tasks SET completed = NOT completed WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
