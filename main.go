package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"go-user-service/database"
	"go-user-service/handler/user"
)

func initDatabase() {
	var err error
	database.DBConnection, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	// // Migrate the schema
	// database.DBConnection.AutoMigrate(&Product{})
	// fmt.Println("Database migrated.")

}

func main() {

	initDatabase()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "User Service.")
	})
	e.GET("/user", user.RegisterUser)

	e.Logger.Fatal(e.Start(":1323"))

	defer database.DBConnection.Close()
}
