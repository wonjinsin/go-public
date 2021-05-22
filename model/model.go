package model

import (
	"context"
	"gorilla/config"
	"gorilla/structs"
	"gorilla/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Logger *utils.Logger

func MongoConn(Gorilla *config.ViperConfig) (db *mongo.Client) {
	credential := options.Credential{
		Username: Gorilla.GetString("db_user"),
		Password: Gorilla.GetString("db_pass"),
	}
	applyUri := Gorilla.GetString("db_host") + ":" + Gorilla.GetString("db_port")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://" + applyUri).SetAuth(credential)
	db, err := mongo.Connect(ctx, clientOptions)
	defer cancel()

	if err != nil {
		Logger.Logging().Errorw("MongoDB Connection Failed")
	}

	// Check the connection
	err = db.Ping(context.TODO(), nil)

	if err != nil {
		Logger.Logging().Errorw("MongoDB check ping failed")
	}

	Logger.Logging().Infow("MongoDB Connection Made")
	return db
}

func ctxGenerator() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	return ctx, cancel
}

type Room interface {
	SetRoomNo(int)
	CheckRoom() (structs.RoomInfo, error)
	GetRoomContents(ctx context.Context) ([]structs.RoomContents, error)
	CreateRoom(ctx context.Context) error
	SendMessage(ctx context.Context) error
	DeleteMessage(ctx context.Context) error
	JoinTest(roomNo int)
}
