package model

import (
	"context"
	"fmt"
	"gorilla/structs"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomModel struct {
	db   *mongo.Client
	room *mongo.Collection
}

func NewRoomModel(db *mongo.Client) *RoomModel {

	room := db.Database("gorilla").Collection("room")

	rm := &RoomModel{
		db:   db,
		room: room,
	}

	return rm
}

func (rm *RoomModel) CheckRoom(c context.Context) (structs.RoomInfo, error) {

	result := structs.RoomInfo{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := rm.room.FindOne(ctx, bson.D{{Key: "roomNo", Value: "1"}}).Decode(&result)

	if err != nil {
		fmt.Println(err)
	}

	return result, err
}
