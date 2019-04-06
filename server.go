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

	e.GET("/", indexHandler)
	e.GET("/tasks", getTasks)
    e.PUT("/tasks", putTask)
    e.DELETE("/tasks/:id", deleteTask)

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

func indexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to GoTask!")
}

func getTasks(c echo.Context) error {
	return c.JSON(200, "GET Tasks")
}

func putTask(c echo.Context) error {
	return c.JSON(200, "PUT Tasks")
}

func deleteTask(c echo.Context) error {
	return c.JSON(200, "DELETE Task "+c.Param("id"))
}

