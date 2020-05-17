package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	// "go-user-service/model"
)

// Handler
func RegisterUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
