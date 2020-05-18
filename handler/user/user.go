package handler

import (
	"net/http"

	"go-user-service/database"

	"github.com/labstack/echo/v4"

	model "go-user-service/model/user"
)

// RegisterUser blah blah some comment
func RegisterUser(c echo.Context) error {

	db := database.DBConnection
	// Create
	db.Create(&model.User{Code: "L1212", Price: 1000})

	// Read
	var user model.User
	db.First(&user, 1)                   // find user with id 1
	db.First(&user, "code = ?", "L1212") // find user with code l1212

	return c.JSON(http.StatusOK, user)
}
