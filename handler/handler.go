package handler

import (
	"context"
	"gorilla/structs"
	"gorilla/utils"
)

var Logger *utils.Logger

type Room interface {
	GetRoom(ctx context.Context) (structs.RoomInfo, error)
	SendMessage(ctx context.Context) error
}
