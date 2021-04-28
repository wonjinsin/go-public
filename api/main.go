package main

import (
	"fmt"
	"gorilla/api/controller"
	"gorilla/config"
	"gorilla/model"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	startServer()
}

func startServer() {
	Gorilla := config.Gorilla
	e := echo.New()
	e.GET("/healthCheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "It's working!")
	})

	db := model.MongoConn(Gorilla)
	controller.InitHandler(Gorilla, e, db)

	e.Start(fmt.Sprintf(":%s", Gorilla.GetString("port")))
}
