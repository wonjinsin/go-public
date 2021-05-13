package model

import (
	"context"
	"fmt"
	"gorilla/structs"
	"gorilla/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (rm *RoomModel) SetRoomNo(ctx context.Context) {
	rm.roomNo = ctx.Value(utils.IntKey(1)).(int)
}

func (rm *RoomModel) CheckRoom() (structs.RoomInfo, error) {
	result := structs.RoomInfo{}

	tmpCtx, cancel := ctxGenerator()
	defer cancel()

	err := rm.room.FindOne(tmpCtx, bson.D{{Key: "roomNo", Value: rm.roomNo}}).Decode(&result)

	if err != nil {
		Logger.Logging().Warnw("No room basic info", "result", err)
		return result, err
	}

	Logger.Logging().Infow("Got room basic info", "result", result)

	return result, err
}

func (rm *RoomModel) GetRoomContents(ctx context.Context) ([]structs.RoomContents, error) {
	result := []structs.RoomContents{}

	tmpCtx, cancel := ctxGenerator()
	defer cancel()

	cur, err := rm.room_contents.Find(tmpCtx, bson.D{{Key: "roomNo", Value: rm.roomNo}})

	if err != nil {
		Logger.Logging().Warnw("No roomContents", "result", err)
		return result, err
	}

	defer cur.Close(tmpCtx)
	for cur.Next(tmpCtx) {
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

	rm.JoinTest(rm.roomNo)

	return result, err
}

func (rm *RoomModel) JoinTest(roomNo int) {
	tmpCtx, cancel := ctxGenerator()
	defer cancel()

	lookupStage := bson.D{{"$lookup", bson.D{{"from", "user"}, {"localField", "user"}, {"foreignField", "name"}, {"as", "podcast"}}}}

	cur, err := rm.room_contents.Aggregate(tmpCtx, lookupStage)

	fmt.Println("------------------------")
	fmt.Println(cur)
	fmt.Println(err)

	for cur.Next(tmpCtx) {
		var row []bson.M
		err := cur.Decode(&row)
		if err != nil {
			Logger.Logging().Warnw("Can't decode result", "result", err)
		}

		fmt.Println(row)
	}

	fmt.Println("------------------------")
}

func (rm *RoomModel) GetRecentOne() (structs.RoomContents, error) {
	result := structs.RoomContents{}

	tmpCtx, cancel := ctxGenerator()
	defer cancel()

	opts := options.FindOne().SetSort(bson.M{"Date": -1})
	err := rm.room_contents.FindOne(tmpCtx, bson.D{{Key: "roomNo", Value: rm.roomNo}}, opts).Decode(&result)

	if err != nil {
		Logger.Logging().Warnw("No recent one", "result", err)
		return result, err
	}

	Logger.Logging().Infow("Got recent one", "result", result)

	return result, err
}

func (rm *RoomModel) SendMessage(ctx context.Context) error {
	roomSendInfo := ctx.Value(utils.StringKey("sendInfo"))
	_, err := rm.room_contents.InsertOne(ctx, roomSendInfo)

	if err != nil {
		Logger.Logging().Warnw("Can't insert message", "result", err)
	}

	return err
}
