package database

import (
	"database/sql"
	"fmt"
	"github.com/jesper.norgard/Todo_go/model"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

const (
	PGHOST      = "PGHOST"
	PGPORT      = "PGPORT"
	PGUSER      = "PGUSER"
	PGDATABASE  = "PGDATABASE"
	PGPASSWORD  = "PGPASSWORD"
	defaultHost = "localhost"
	defaultPort = 5432
	defaultUser = "jesper.norgard"
	defaultDB   = "go_test_db"
	defaultPW   = "."
)

func SetupDb() *sql.DB {

	host := getEnv(PGHOST, defaultHost)
	port := getEnvAsInt(PGPORT, defaultPort)
	user := getEnv(PGUSER, defaultUser)
	dbname := getEnv(PGDATABASE, defaultDB)
	password := getEnv(PGPASSWORD, defaultPW)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	createTables(db)
	return db
}

func createTables(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS tasks (id SERIAL PRIMARY KEY , name TEXT NOT NULL ,completed BOOLEAN\n)")
	if err != nil {
		panic(err)
	}
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

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
