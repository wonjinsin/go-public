package handler

import (
	"context"
	"giraffe/config"
	"giraffe/model"
	"giraffe/structs"
	"giraffe/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rh *RoomHandler

func TestMain(m *testing.M) {
	Giraffe := config.Giraffe

	db := model.MongoConn(Giraffe)
	rh = NewRoomHandler(db)
	os.Exit(m.Run())
}

func TestGetRoom(t *testing.T) {
	assert := assert.New(t)

	var key utils.IntKey = 1
	roomNo := 1
	ctx, _ := utils.CtxGenerator()
	ctx = context.WithValue(ctx, key, roomNo)

	_, err := rh.GetRoom(ctx)

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
}

func TestCreateRoom(t *testing.T) {
	assert := assert.New(t)

	var key utils.StringKey = "roomCreateInfo"
	roomCreateInfo := structs.RoomCreateInfo{RoomNo: 15, Users: [2]string{"wonjinsin", "wonjinsin2"}}
	ctx, _ := utils.CtxGenerator()
	ctx = context.WithValue(ctx, key, roomCreateInfo)

	err := rh.CreateRoom(ctx)

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
}

func TestSendMessage(t *testing.T) {
	assert := assert.New(t)

	var key utils.StringKey = "roomSendInfo"
	roomSendInfo := structs.RoomSendInfo{RoomNo: 15, User: "wonjinsin", Message: "Test messgage Send", Date: utils.GetNow()}

	ctx, _ := utils.CtxGenerator()
	ctx = context.WithValue(ctx, key, roomSendInfo)

	err := rh.SendMessage(ctx)

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
}

func TestDeleteMessage(t *testing.T) {
	assert := assert.New(t)

	var key utils.StringKey = "roomDeleteInfo"
	roomDeleteInfo := structs.RoomDeleteInfo{
		ObjectId: "60cdd8399cb62167970f1451",
	}
	ctx, _ := utils.CtxGenerator()
	ctx = context.WithValue(ctx, key, roomDeleteInfo)

	err := rh.DeleteMessage(ctx)

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
}
