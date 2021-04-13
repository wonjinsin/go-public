package main

import (
	"chatapp/api/controller"
	"chatapp/config"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	startServer()
}

func startServer() {
	Chatapp := config.Chatapp
	e := echo.New()
	e.GET("/healthCheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "It's working!")
	})

	controller.InitHandler(e)
	e.Start(fmt.Sprintf(":%s", Chatapp.GetString("port")))
}
