package handler

import (
	"context"
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

func (uh *UserHandler) Login(ctx context.Context) error {
	err := uh.md.GetUser(ctx)

	if err != nil {
		Logger.Logging().Warnw("Fail to Login", "result", err)
		return err
	}

	return err
}
