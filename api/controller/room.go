package controller

import (
	"gorilla/config"
	"gorilla/handler"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type httpRoomController struct {
	db *mongo.Client
	rh *handler.RoomHandler
}

func newHTTPRoomContoller(gorilla *config.ViperConfig, eg *echo.Group, db *mongo.Client) {
	rh := handler.NewRoomHandler(db)

	h := &httpRoomController{
		db: db,
		rh: rh,
	}

	eg.GET("", h.Room)
}

func (h *httpRoomController) Room(c echo.Context) error {
	ctx := c.Request().Context()
	h.rh.GetRoom(ctx)

	return response(c, 200, "test message", "test")
}
