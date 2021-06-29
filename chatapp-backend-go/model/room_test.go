package model

import (
	"context"
	"giraffe/config"
	"giraffe/structs"
	"giraffe/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client
var rm *RoomModel

func TestMain(m *testing.M) {
	Giraffe := config.Giraffe

	db = MongoConn(Giraffe)
	rm = NewRoomModel(db)
	os.Exit(m.Run())
}

func TestCheckRoom(t *testing.T) {
	assert := assert.New(t)

	rm.SetRoomNo(13)
	_, err := rm.CheckRoom()

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
}

func TestGetRoomContents(t *testing.T) {
	assert := assert.New(t)

	rm.SetRoomNo(13)
	_, err := rm.GetRoomContents()

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
}

func TestGetRecentOne(t *testing.T) {
	assert := assert.New(t)

	rm.SetRoomNo(13)
	_, err := rm.GetRecentOne()

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

	err := rm.CreateRoom(ctx)

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
}

func TestSendMessage(t *testing.T) {
	assert := assert.New(t)

	var key utils.StringKey = "roomSendInfo"
	roomSendInfo := structs.RoomSendInfo{RoomNo: 15, User: "wonjinsin", Message: "Test messgage Send Test!", Date: utils.GetNow()}

	ctx, _ := utils.CtxGenerator()
	ctx = context.WithValue(ctx, key, roomSendInfo)

	err := rm.SendMessage(ctx)

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
}

func TestDeleteRoom(t *testing.T) {
	assert := assert.New(t)

	var key utils.StringKey = "roomDeleteInfo"
	roomDeleteInfo := structs.RoomDeleteInfo{
		ObjectId: "60cdd4cdf181cb604c78b7fd",
	}
	ctx, _ := utils.CtxGenerator()
	ctx = context.WithValue(ctx, key, roomDeleteInfo)

	err := rm.DeleteRoom(ctx)

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
}

func TestDeleteMessage(t *testing.T) {
	assert := assert.New(t)

	var key utils.StringKey = "messageDeleteInfo"
	roomDeleteInfo := structs.RoomDeleteInfo{
		ObjectId: "60cf2854ec5adece8c2b89e0",
	}
	ctx, _ := utils.CtxGenerator()
	ctx = context.WithValue(ctx, key, roomDeleteInfo)

	err := rm.DeleteMessage(ctx)

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
}
