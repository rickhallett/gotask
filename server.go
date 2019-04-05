package main

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo"
)

// User to provide bind json payload into Go struct based on Content-Type JSON request header
type User struct {
	Fname string `json:"firstname"`
	Lname string `json:"lastname"`
	Username string `json:"username"`
	Email string `json:"email"`
}

// UserDB can act as a mock user DB table until one is created
var UserDB []User

func main() {
	e := echo.New()
	e.GET("/", indexHandler)
	e.POST("/register", register)

	e.Logger.Fatal(e.Start(":1323"))
}

func indexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to GoTask!")
}

func register(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}


