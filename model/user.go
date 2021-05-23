package model

import (
	"context"
	"fmt"
	"gorilla/structs"
	"gorilla/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	db   *mongo.Client
	user *mongo.Collection
}

func NewUserModel(db *mongo.Client) *UserModel {

	user := db.Database("gorilla").Collection("user")

	um := &UserModel{
		db:   db,
		user: user,
	}

	return um
}

func (um *UserModel) GetUser(ctx context.Context) error {
	tmpCtx, cancel := ctxGenerator()
	defer cancel()

	result := structs.User{}
	loginInfo := ctx.Value(utils.StringKey("loginInfo"))
	fmt.Println(loginInfo)

	// need fix
	err := um.user.FindOne(tmpCtx, loginInfo).Decode(&result)

	if err != nil {
		Logger.Logging().Warnw("User not exist", "result", err)
	}

	return err
}
