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
	// result, err := h.rh.GetRoom(ctx)
	result := h.rh.GetRoom(ctx)

	return response(c, 404, "Room is not exist", result)
	//	if err != nil {
	//		return response(c, 404, "Room is not exist", "")
	//	}
	//
	//	return response(c, 200, "Got room Info", result)
}
