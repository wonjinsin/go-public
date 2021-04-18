package controller

import (
	"gorilla/config"

	"github.com/labstack/echo"
)

type httpRoomHandler struct {
}

type JsonInterface struct {
	Test string `json:"test"`
}

func newHTTPRoomHandler(Gorilla *config.ViperConfig, eg *echo.Group) {
	h := &httpRoomHandler{}
	eg.GET("", h.Room)
}

func (h *httpRoomHandler) Room(c echo.Context) error {
	roomJson := JsonInterface{
		Test: "Test",
	}

	return response(c, 200, "test message", roomJson)
}
