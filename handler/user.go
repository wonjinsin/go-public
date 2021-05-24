package handler

import (
	"context"
	"gorilla/model"

	"github.com/dgrijalva/jwt-go"
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
	result, err := uh.md.GetUser(ctx)

	if err != nil {
		Logger.Logging().Warnw("Fail to Login", "result", err)
		return err
	}

	return err
}

func createToken(user string) (string, error) {
	// create a signer for rsa 256
	jwt.New(jwt.GetSigningMethod("RS256"))
}
