package handler

import (
	"context"
	"fmt"
	"gorilla/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	db *mongo.Client
	md model.User
}

func NewUserHandler(db *mongo.Client) *UserHandler {
	md := model.NewUserModel(db)

	uh := &UserHandler{
		db: db,
		md: md,
	}

	return uh
}

func (uh *UserHandler) Login(ctx context.Context) (string, error) {
	var token string
	result, err := uh.md.GetUser(ctx)

	if err != nil {
		Logger.Logging().Warnw("Fail to Login", "result", err)
		return token, err
	}

	tk, err := NewTokenHandler()

	if err != nil {
		Logger.Logging().Warnw("Fail to init tokenHandler", "result", err)
		return token, err
	}

	token, err = tk.createToken(result)

	if err != nil {
		Logger.Logging().Warnw("Fail to createToken", "result", err)
		return token, err
	}

	fmt.Println(token)
	return token, err
}
