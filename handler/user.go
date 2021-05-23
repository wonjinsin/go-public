package handler

import (
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
