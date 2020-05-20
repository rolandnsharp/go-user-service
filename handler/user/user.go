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
	db.Create(&model.User{Name: "L1212", Email: "blah@email.com"})

	var user model.User
	db.First(&user, 1)                   // find user with id 1
	db.First(&user, "name = ?", "L1212") // find user with code l1212

	return c.JSON(http.StatusOK, user)
}
