package handler

import (
	"context"
	"gorilla/structs"
	"gorilla/utils"

	"github.com/labstack/echo"
)

var Logger *utils.Logger

type Room interface {
	GetRoom(ctx context.Context) (structs.RoomInfo, error)
	CreateRoom(ctx context.Context) error
	SendMessage(ctx context.Context) error
	DeleteMessage(ctx context.Context) error
}

type User interface {
	Login(ctx context.Context) (string, error)
	Validate(c echo.Context) (structs.User, error)
}

type Token interface {
	CreateToken(user structs.User) (string, error)
	ValidateToken(c echo.Context) (structs.User, error)
}
