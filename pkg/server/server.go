package server

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func shutdown() {
	os.Exit(0)
}
