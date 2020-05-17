package user

import (
	"fmt"
	"net/http"

	"go-user-service/database"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func RegisterUser(c echo.Context) error {

	db := database.DBConnection
	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	fmt.Println(product)
	return c.String(http.StatusOK, "Hello, World!")
}
