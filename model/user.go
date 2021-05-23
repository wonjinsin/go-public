package model

import "go.mongodb.org/mongo-driver/mongo"

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
