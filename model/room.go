package model

import (
	"context"
	"gorilla/structs"
	"gorilla/utils"
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
		Logger.Logging().Warnw("No room basic info", "result", err)
		return result, err
	}

	Logger.Logging().Infow("Got room basic info", "result", result)

	return result, err
}

func (rm *RoomModel) GetRoomContents(c context.Context) ([]structs.RoomContents, error) {
	result := []structs.RoomContents{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := rm.room_contents.Find(ctx, bson.D{{Key: "roomNo", Value: rm.roomNo}})

	if err != nil {
		Logger.Logging().Warnw("No roomContents", "result", err)
		return result, err
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var row structs.RoomContents
		err := cur.Decode(&row)
		if err != nil {
			Logger.Logging().Warnw("Can't decode result", "result", err)
		}

		row.DateStr = utils.TimeFormat(row.Date)
		result = append(result, row)
	}

	if err := cur.Err(); err != nil {
		Logger.Logging().Warnw("Can't decode result", "result", err)
	}

	Logger.Logging().Infow("Got roomContents", "result", result)

	return result, err
}
