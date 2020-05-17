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

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func initDatabase() {
	var err error
	database.DBConnection, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	// // Migrate the schema
	database.DBConnection.AutoMigrate(&Product{})
}

func main() {
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// // Migrate the schema
	// db.AutoMigrate(&Product{})
	// // Create
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// // Read
	// var product Product
	// db.First(&product, 1)                   // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// fmt.Println(product)

	initDatabase()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "User Service.")
	})
	e.GET("/user", user.RegisterUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

	defer database.DBConnection.Close()
}
