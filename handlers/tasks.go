package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	// "gotask/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "GET Tasks")
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, H{
			"created": "yup",
		})
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        return c.JSON(http.StatusOK, H{
            "deleted": id,
        })
    }
}
