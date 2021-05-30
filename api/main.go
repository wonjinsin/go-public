package main

import (
	"fmt"
	"giraffe/api/controller"
	"giraffe/config"
	"giraffe/model"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	startServer()
}

func startServer() {
	Giraffe := config.Giraffe
	e := echo.New()
	e.GET("/healthCheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "It's working!")
	})

	db := model.MongoConn(Giraffe)
	controller.InitHandler(Giraffe, e, db)

	e.Start(fmt.Sprintf(":%s", Giraffe.GetString("port")))
}
