package model

import (
	"context"
	"giraffe/structs"
	"giraffe/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	db   *mongo.Client
	user *mongo.Collection
}

func NewUserModel(db *mongo.Client) *UserModel {

	user := db.Database("giraffe").Collection("user")

	um := &UserModel{
		db:   db,
		user: user,
	}

	return um
}

func (um *UserModel) GetUser(ctx context.Context) (structs.User, error) {
	tmpCtx, cancel := CtxGenerator()
	defer cancel()

	result := structs.User{}
	loginInfo := ctx.Value(utils.StringKey("loginInfo"))

	// need fix
	err := um.user.FindOne(tmpCtx, loginInfo).Decode(&result)

	if err != nil {
		Logger.Logging().Warnw("User not exist", "result", err)
	}

	return result, err
}
