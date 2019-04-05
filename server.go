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

// DB can act as a mock DB until one is created
type DB struct{
	User []*User
}

// VDB Virtual Database
var VDB DB

func (db *DB) initDB(){
	VDB.User = []*User{}
}

func (db *DB) insertUser(u *User) {
	db.User = append(db.User, u)
	for i, u := range db.User {
		fmt.Printf("id: %v, user: %+v\n", i, u)
	}
}

// Mock database global storage
func init(){
	VDB = DB{}
	VDB.initDB()
}

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
	
	VDB.insertUser(u)
	return c.JSON(http.StatusCreated, u)
}



