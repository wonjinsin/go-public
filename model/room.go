package model

import (
	"context"
	"fmt"
	"gorilla/structs"
	"gorilla/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomModel struct {
	db            *mongo.Client
	room          *mongo.Collection
	room_contents *mongo.Collection
	roomNo        int
}

func NewRoomModel(db *mongo.Client) *RoomModel {

	room := db.Database("gorilla").Collection("room")
	room_contents := db.Database("gorilla").Collection("room_contents")

	rm := &RoomModel{
		db:            db,
		room:          room,
		room_contents: room_contents,
	}

	return rm
}

func (rm *RoomModel) SetRoomNo(c context.Context) {
	rm.roomNo = c.Value(utils.IntKey(1)).(int)
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
	result := []structs.RoomContents{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := rm.room_contents.Find(ctx, bson.D{{Key: "roomNo", Value: rm.roomNo}})

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var row structs.RoomContents
		err := cur.Decode(&row)
		if err != nil {
			log.Fatal(err)
		}

		row.DateStr = utils.TimeFormat(row.Date)
		result = append(result, row)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return result, err
}
