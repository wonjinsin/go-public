package handler

import (
	"context"
	"giraffe/config"
	"giraffe/model"
	"giraffe/structs"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	db      *mongo.Client
	md      model.User
	tk      Token
	giraffe *config.ViperConfig
}

func NewUserHandler(db *mongo.Client, giraffe *config.ViperConfig) *UserHandler {
	md := model.NewUserModel(db)
	uh := &UserHandler{
		db:      db,
		md:      md,
		giraffe: giraffe,
	}

	tk, err := NewTokenHandler(uh.giraffe)

	if err != nil {
		Logger.Logging().Errorw("Fail to init tokenHandler", "result", err)
	}

	uh.tk = tk

	return uh
}

func (uh *UserHandler) Login(ctx context.Context) (string, error) {
	var token string
	result, err := uh.md.GetUser(ctx)

	if err != nil {
		Logger.Logging().Warnw("Fail to Login", "result", err)
		return token, err
	}

	token, err = uh.tk.CreateToken(result)

	if err != nil {
		Logger.Logging().Warnw("Fail to createToken", "result", err)
		return token, err
	}

	return token, err
}

func (uh *UserHandler) Validate(c echo.Context) (structs.User, error) {
	user, err := uh.tk.ValidateToken(c)

	if err != nil {
		Logger.Logging().Warnw("Fail to ValidateUserToken", "result", err)
		return structs.User{}, err
	}

	return user, nil
}
