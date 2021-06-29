package controller

import (
	"context"
	"giraffe/config"
	"giraffe/handler"
	"giraffe/structs"
	"giraffe/utils"
	"strconv"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type httpRoomController struct {
	db *mongo.Client
	rh handler.Room
}

func newHTTPRoomContoller(giraffe *config.ViperConfig, eg *echo.Group, db *mongo.Client) {
	rh := handler.NewRoomHandler(db)

	h := &httpRoomController{
		db: db,
		rh: rh,
	}

	eg.GET("/:roomNo", h.Room)
	eg.POST("/create", h.Create)
	eg.POST("/send", h.Send)
	eg.DELETE("/delete/:roomNo", h.DeleteRoom)
	eg.DELETE("/message/:objectId", h.DeleteMessage)
}

func (h *httpRoomController) Room(c echo.Context) error {
	roomNoStr := c.Param("roomNo")
	roomNo, err := strconv.Atoi(roomNoStr)

	if err != nil {
		return response(c, 404, "Parameter should be number", utils.ErrorToStr(err))
	}

	var key utils.IntKey = 1
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, key, roomNo)

	result, err := h.rh.GetRoom(ctx)

	if err != nil {
		return response(c, 404, "Room is not exist", utils.ErrorToStr(err))
	}

	return response(c, 200, "Got room Info", result)
}

func (h *httpRoomController) Create(c echo.Context) error {

	roomCreate := &structs.RoomCreate{}
	err := c.Bind(roomCreate)

	if err != nil {
		return response(c, 404, "format is not valid")
	}

	roomNoStr := roomCreate.RoomNo
	roomNo, err := strconv.Atoi(roomNoStr)

	if err != nil {
		return response(c, 404, "Parameter should be number")
	}

	user1 := roomCreate.User1
	user2 := roomCreate.User2

	if user1 == "" || user2 == "" {
		return response(c, 404, "user is not exist")
	}

	roomCreateInfo := structs.RoomCreateInfo{}
	roomCreateInfo.RoomNo = roomNo
	roomCreateInfo.Users = [2]string{user1, user2}

	var key utils.StringKey = "roomCreateInfo"
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, key, roomCreateInfo)

	err = h.rh.CreateRoom(ctx)

	if err != nil {
		return response(c, 404, "Create room failed", utils.ErrorToStr(err))
	}

	return response(c, 200, "Create room succeeded", roomCreateInfo)
}

func (h *httpRoomController) Send(c echo.Context) error {

	roomSend := &structs.RoomSend{}
	c.Bind(roomSend)

	roomNoStr := roomSend.RoomNo
	roomNo, err := strconv.Atoi(roomNoStr)

	if err != nil {
		return response(c, 404, "Parameter should be number", "")
	}

	roomSendInfo := &structs.RoomSendInfo{}
	roomSendInfo.RoomNo = roomNo
	roomSendInfo.User = roomSend.User
	roomSendInfo.Message = roomSend.Message
	roomSendInfo.Date = utils.GetNow()

	var key utils.StringKey = "roomSendInfo"
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, key, roomSendInfo)

	err = h.rh.SendMessage(ctx)

	if err != nil {
		return response(c, 404, "Failed to insert message", utils.ErrorToStr(err))
	}

	return response(c, 200, "Success insert message", roomSendInfo)
}

func (h *httpRoomController) DeleteRoom(c echo.Context) error {
	roomNoStr := c.Param("roomNo")
	roomNo, err := strconv.Atoi(roomNoStr)

	if err != nil {
		return response(c, 404, "Parameter should be a number")
	}

	var key utils.IntKey = 1
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, key, roomNo)

	err = h.rh.DeleteRoom(ctx)

	if err != nil {
		return response(c, 404, "Failed to delete room")
	}

	return response(c, 200, "Success DeleteRoom")
}

func (h *httpRoomController) DeleteMessage(c echo.Context) error {
	objectId := c.Param("objectId")

	obj := structs.RoomDeleteInfo{
		ObjectId: objectId,
	}

	var key utils.StringKey = "messageDeleteInfo"
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, key, obj)

	err := h.rh.DeleteMessage(ctx)

	if err != nil {
		return response(c, 405, "Delete message failed", utils.ErrorToStr(err))
	}

	return response(c, 200, "Success insert message", obj)
}
