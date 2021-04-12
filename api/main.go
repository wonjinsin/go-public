package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	startServer()
}

func startServer() {

	e := echo.New()
	e.GET("/healthCheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "It's working!")
	})

	e.Start(":12001")

}
