package handler

import "context"

type Room interface {
	CheckRoom(ctx context.Context, room string) (err error)
	GetRoomContents(ctx context.Context, room string) (err error)
}
