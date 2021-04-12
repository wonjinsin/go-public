package controller

import "github.com/labstack/echo"

func initHandler(e *echo.Echo) {
	api := e.Group("/api")
	ver1 := api.Group("/v1")

	room := ver1.Group("/room")

}
