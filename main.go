package main

import (
	"belajar_rest_api/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Assalaam ")
	})

	routes.StudentRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
