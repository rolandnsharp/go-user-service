package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-user-service/handler/user"
)

func main() {
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
}
