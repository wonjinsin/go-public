package handler

import (
	"context"
	"gorilla/model"
	"gorilla/structs"

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
	rh.md.SetRoomNo(ctx)
	roomInfo, err := rh.md.CheckRoom()

	if err != nil {
		return roomInfo, err
	}

	roomContents, err := rh.md.GetRoomContents(ctx)
	roomInfo.Contents = roomContents

	return roomInfo, err
}
