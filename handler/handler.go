package handler

import (
	"context"
	"gorilla/structs"
	"gorilla/utils"
)

var Logger *utils.Logger

type Room interface {
	GetRoom(ctx context.Context) (structs.RoomInfo, error)
	CreateRoom(ctx context.Context) error
	SendMessage(ctx context.Context) error
}
