package handler

import (
	"context"
	"errors"
	"giraffe/model"
	"giraffe/structs"
	"giraffe/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type RoomHandler struct {
	db *mongo.Client
	md model.Room
}

func NewRoomHandler(db *mongo.Client) *RoomHandler {
	md := model.NewRoomModel(db)

	rh := &RoomHandler{
		db: db,
		md: md,
	}

	return rh
}

func (rh *RoomHandler) GetRoom(ctx context.Context) (structs.RoomInfo, error) {
	roomNo := ctx.Value(utils.IntKey(1)).(int)
	rh.md.SetRoomNo(roomNo)

	roomInfo, err := rh.md.CheckRoom()

	if err != nil {
		Logger.Logging().Warnw("Got roomInfo error", "result", err)
		return roomInfo, err
	}

	roomContents, err := rh.md.GetRoomContents(ctx)
	if err != nil {
		Logger.Logging().Warnw("Got roomContents error", "result", err)
		return roomInfo, err
	}

	roomInfo.Contents = roomContents
	Logger.Logging().Infow("Got roomInfo", "result", roomInfo)

	return roomInfo, err
}

func (rh *RoomHandler) CreateRoom(ctx context.Context) error {
	roomNo := ctx.Value(utils.StringKey("roomCreateInfo")).(structs.RoomCreateInfo).RoomNo

	rh.md.SetRoomNo(roomNo)

	result, err := rh.md.CheckRoom()

	if err == nil {
		Logger.Logging().Warnw("Room is exist", "result", result)
		return errors.New("Room is exist")
	}

	err = rh.md.CreateRoom(ctx)

	if err != nil {
		Logger.Logging().Warnw("Create room Failed", "result", err)
	}

	return err
}

func (rh *RoomHandler) SendMessage(ctx context.Context) error {
	err := rh.md.SendMessage(ctx)

	if err != nil {
		Logger.Logging().Warnw("Insert message error", "result", err)
	}

	return err
}

func (rh *RoomHandler) DeleteRoom(ctx context.Context) error {
	return errors.New("test")
}

func (rh *RoomHandler) DeleteMessage(ctx context.Context) error {
	err := rh.md.DeleteMessage(ctx)

	if err != nil {
		Logger.Logging().Warnw("Delete message failed", "result", err)
	}

	return err
}
