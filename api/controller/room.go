package controller

import (
	"context"
	"gorilla/config"
	"gorilla/handler"
	"gorilla/structs"
	"gorilla/utils"
	"strconv"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type httpRoomController struct {
	db *mongo.Client
	rh handler.Room
}

func newHTTPRoomContoller(gorilla *config.ViperConfig, eg *echo.Group, db *mongo.Client) {
	rh := handler.NewRoomHandler(db)

	h := &httpRoomController{
		db: db,
		rh: rh,
	}

	eg.GET("/:roomNo", h.Room)
	eg.POST("/send", h.Send)
}

func (h *httpRoomController) Room(c echo.Context) error {
	roomNoStr := c.Param("roomNo")
	roomNo, err := strconv.Atoi(roomNoStr)

	if err != nil {
		return response(c, 404, "Parameter should be number", "")
	}

	var key utils.IntKey = 1
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, key, roomNo)

	result, err := h.rh.GetRoom(ctx)

	if err != nil {
		return response(c, 404, "Room is not exist", "")
	}

	return response(c, 200, "Got room Info", result)
}

func (h *httpRoomController) Create(c echo.Context) structs.RoomContents {
	roomContents := structs.RoomContents{}
	roomContents.User = c.FormValue("User")
	roomContents.Message = c.FormValue("Message")

	return roomContents
}
