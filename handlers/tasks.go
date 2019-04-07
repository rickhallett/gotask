package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"fmt"

	"gotask/models"

	"github.com/labstack/echo"
)

// H ...
type H map[string]interface{}

// GetTasks ...
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTask ...
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task models.Task
		c.Bind(&task)

		id, err := models.PutTask(db, task.Name)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("Task %v created: %s\n", id, task.Name)
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		}
		
	}
}

// DeleteTask ...
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// TODO: DeleteTask needs to only return deleted message if a todo that existed was deleted.
		_, err := models.DeleteTask(db, id)
		if err != nil {
			panic(err)
		} else {
			return c.JSON(http.StatusOK, H{
            	"deleted": id,
        	})
		}
        
    }
}
