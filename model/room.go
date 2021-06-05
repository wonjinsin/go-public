package model

import (
	"context"
	"errors"
	"giraffe/structs"
	"giraffe/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	room := db.Database("giraffe").Collection("room")
	room_contents := db.Database("giraffe").Collection("room_contents")

	rm := &RoomModel{
		db:            db,
		room:          room,
		room_contents: room_contents,
	}

	return rm
}

func (rm *RoomModel) SetRoomNo(num int) {
	rm.roomNo = num
}

func (rm *RoomModel) CheckRoom() (structs.RoomInfo, error) {
	result := structs.RoomInfo{}

	tmpCtx, cancel := CtxGenerator()
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

	tmpCtx, cancel := CtxGenerator()
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

	return result, err
}

func (rm *RoomModel) JoinTest(roomNo int) {
	tmpCtx, cancel := CtxGenerator()
	defer cancel()

	project := bson.M{
		"$lookup": bson.M{
			"from":         "user",
			"localField":   "user",
			"foreignField": "name",
			"as":           "test",
		},
	}

	cur, err := rm.room_contents.Aggregate(tmpCtx, []bson.M{
		project,
	})

	Logger.Logging().Warnw("Can't search", "result", err)

	for cur.Next(tmpCtx) {
		row := bson.M{}
		err := cur.Decode(&row)
		if err != nil {
			Logger.Logging().Warnw("Can't decode result", "result", err)
		}
	}

}

func (rm *RoomModel) GetRecentOne() (structs.RoomContents, error) {
	result := structs.RoomContents{}

	tmpCtx, cancel := CtxGenerator()
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

func (rm *RoomModel) CreateRoom(ctx context.Context) error {
	tmpCtx, cancel := CtxGenerator()
	defer cancel()

	createRoomInfo := ctx.Value(utils.StringKey("roomCreateInfo"))
	_, err := rm.room.InsertOne(tmpCtx, createRoomInfo)

	if err != nil {
		Logger.Logging().Warnw("Can't create room", "result", err)
	}

	return err
}

func (rm *RoomModel) SendMessage(ctx context.Context) error {
	tmpCtx, cancel := CtxGenerator()
	defer cancel()

	roomSendInfo := ctx.Value(utils.StringKey("roomSendInfo"))
	_, err := rm.room_contents.InsertOne(tmpCtx, roomSendInfo)

	if err != nil {
		Logger.Logging().Warnw("Can't insert message", "result", err)
	}

	return err
}

func (rm *RoomModel) DeleteMessage(ctx context.Context) error {
	tmpCtx, cancel := CtxGenerator()
	defer cancel()

	objectId := ctx.Value(utils.StringKey("roomDeleteInfo")).(structs.RoomDeleteInfo).ObjectId
	objectIdHex, err := primitive.ObjectIDFromHex(objectId)
	errMsg := "Can't delete message"
	noResultMsg := "No result found to delete"

	if err != nil {
		Logger.Logging().Warnw(errMsg, "result", err)
		return err
	}

	result, err := rm.room_contents.DeleteOne(tmpCtx, bson.M{"_id": objectIdHex})

	if result.DeletedCount == 0 {
		Logger.Logging().Warnw(noResultMsg)
		return errors.New(noResultMsg + " " + objectId)
	}

	if err != nil {
		Logger.Logging().Warnw(errMsg, "result", err)
	}

	return err
}
