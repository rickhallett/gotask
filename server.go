package main

import (
	"database/sql"
	"net/http"

	"gotask/handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()

	db := initDB("gotask.db")
	migrate(db)

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
    e.PUT("/tasks", handlers.putTask(db))
    e.DELETE("/tasks/:id", handlers.deleteTask(db))

	e.Logger.Fatal(e.Start(":8888"))

}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL
	);
	`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

