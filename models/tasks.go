package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Task ...
type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

// TaskCollection ...
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks ...
func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		scanErr := rows.Scan(&task.ID, &task.Name)
		if scanErr != nil {
			panic(scanErr)
		}
		result.Tasks = append(result.Tasks, task)
	}

	return result
}

// PutTask ...
func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, execErr := stmt.Exec(name)
	if execErr != nil {
		panic(execErr)
	}

	return result.LastInsertId()
}

// DeleteTask ...
func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, execErr := stmt.Exec(id)
	if execErr != nil {
		panic(execErr)
	}

	return result.RowsAffected()
}
