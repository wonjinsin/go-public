package model

import (
	"context"
	"fmt"
	"log"
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

func (rm *RoomModel) CheckRoom(ctx context.Context) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := rm.room.Find(ctx, bson.D{{Key: "roomNumber", Value: "1"}})

	if err != nil {
		fmt.Println(err)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
		// do something with result....
	}
	return true
}
