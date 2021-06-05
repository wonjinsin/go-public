package handler

import (
	"context"
	"giraffe/config"
	"giraffe/model"
	"giraffe/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client

func TestMain(m *testing.M) {
	Giraffe := config.Giraffe

	db = model.MongoConn(Giraffe)
	os.Exit(m.Run())
}

func TestGetRoom(t *testing.T) {
	assert := assert.New(t)

	var key utils.IntKey = 1
	ctx, _ := model.CtxGenerator()
	ctx = context.WithValue(ctx, key, 1)

	rh := NewRoomHandler(db)
	_, err := rh.GetRoom(ctx)

	assert.Error(err)
}
