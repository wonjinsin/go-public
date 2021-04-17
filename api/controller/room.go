package controller

import (
	"gorilla/config"

	"github.com/labstack/echo"
)

type httpRoomHandler struct {
}

func newHTTPRoomHandler(Gorilla *config.ViperConfig, eg *echo.Group) {
	h := &httpRoomHandler{}
	eg.GET("/", h.Room)
}

func (h *httpRoomHandler) Room(c echo.Context) error {

	roomJson := map[string]string
	roomJson.test = "test"

	return response(c, 200, "test message", roomJson)
}
