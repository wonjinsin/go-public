package handler

import (
	"context"
	"gorilla/model"
	"gorilla/structs"

	"go.mongodb.org/mongo-driver/mongo"
)

type RoomHandler struct {
	db *mongo.Client
	md *model.RoomModel
}

func NewRoomHandler(db *mongo.Client) *RoomHandler {
	md := model.NewRoomModel(db)

	rh := &RoomHandler{
		db: db,
		md: md,
	}

	return rh
}

func (rh *RoomHandler) GetRoom(ctx context.Context) structs.RoomInfo {
	roomInfo, _ := rh.md.CheckRoom(ctx)

	return roomInfo
}
