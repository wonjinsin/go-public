package model

import (
	"context"
	"fmt"
	"gorilla/structs"
	"gorilla/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomModel struct {
	db     *mongo.Client
	room   *mongo.Collection
	roomNo interface{}
}

func NewRoomModel(db *mongo.Client) *RoomModel {

	room := db.Database("gorilla").Collection("room")

	rm := &RoomModel{
		db:   db,
		room: room,
	}

	return rm
}

func (rm *RoomModel) SetRoomNo(c context.Context) {
	rm.roomNo = c.Value(utils.IntKey(1))
}

func (rm *RoomModel) CheckRoom() (structs.RoomInfo, error) {
	result := structs.RoomInfo{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := rm.room.FindOne(ctx, bson.D{{Key: "roomNo", Value: rm.roomNo}}).Decode(&result)

	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

func (rm *RoomModel) GetRoomContents(c context.Context) ([]structs.RoomContents, error) {
	userId := c.Value(utils.IntKey(1))
	result := []structs.RoomContents{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := rm.room.FindOne(ctx, bson.D{{Key: "roomNo", Value: userId}}).Decode(&result)

	if err != nil {
		fmt.Println(err)
	}

	return result, err
}
