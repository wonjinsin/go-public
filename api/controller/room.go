package controller

import (
	"context"
	"gorilla/config"
	"gorilla/handler"
	"gorilla/utils"

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

	eg.GET("/:userId", h.Room)
}

func (h *httpRoomController) Room(c echo.Context) error {
	userId := c.Param("userId")
	var key utils.StringKey = "userId"
	// if err != nil {
	// return response(c, 404, "Param is not number", "")
	// }

	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, key, userId)
	result := h.rh.GetRoom(ctx)

	return response(c, 404, "Room is not exist", result)
	//	if err != nil {
	//		return response(c, 404, "Room is not exist", "")
	//	}
	//
	//	return response(c, 200, "Got room Info", result)
}
