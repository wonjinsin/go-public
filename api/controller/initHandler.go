package controller

import (
	"chatapp/config"

	"github.com/labstack/echo"
)

func InitHandler(Chatapp *config.ViperConfig, e *echo.Echo) {

	// api := e.Group("/api")
	// ver1 := api.Group("/v1")

	// room := ver1.Group("/room")
	// newHTTPRoomHandler(conf, room)

}
