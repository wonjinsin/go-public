package handler

import (
	"context"
	"gorilla/structs"
)

type Room interface {
	GetRoom(ctx context.Context) (structs.RoomInfo, error)
}
