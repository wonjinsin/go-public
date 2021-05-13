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

func (h *httpRoomController) Send(c echo.Context) error {
	roomNoStr := c.FormValue("roomNo")
	roomNo, err := strconv.Atoi(roomNoStr)

	if err != nil {
		return response(c, 404, "Parameter should be number", "")
	}

	roomSendInfo := structs.RoomSendInfo{}
	roomSendInfo.RoomNo = roomNo
	roomSendInfo.User = c.FormValue("user")
	roomSendInfo.Message = c.FormValue("message")
	roomSendInfo.Date = utils.GetNow()

	var key utils.StringKey = "sendInfo"
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, key, roomSendInfo)

	err = h.rh.SendMessage(ctx)

	if err != nil {
		return response(c, 404, "Insert message failed", err)
	}

	return response(c, 200, "Success insert message", roomSendInfo)
}
