package main

import (
	"database/sql"
	"net/http"

	"gotask/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()

	allowCORSConfig := middleware.CORSConfig{
  		AllowOrigins: []string{"http://localhost:8080"},
  		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}

	e.Use(middleware.CORSWithConfig(allowCORSConfig))

	db := initDB("gotask.db")
	migrate(db)

	e.GET("/tasks", handlers.GetTasks(db))
    e.PUT("/tasks", handlers.PutTask(db))
    e.DELETE("/tasks/:id", handlers.DeleteTask(db))

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

