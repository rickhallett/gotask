package main

import (
	"net/http"
	"github.com/labstack/echo"
	// "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	e.GET("/", indexHandler)

	e.GET("/tasks", getTasks)
    e.PUT("/tasks", putTask)
    e.DELETE("/tasks/:id", deleteTask)

	e.Logger.Fatal(e.Start(":8888"))

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

