package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
func RegisterUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
