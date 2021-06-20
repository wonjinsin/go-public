package model

import (
	"context"
	"giraffe/config"
	"giraffe/structs"
	"giraffe/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Logger *utils.Logger

func MongoConn(Giraffe *config.ViperConfig) (db *mongo.Client) {
	credential := options.Credential{
		Username: Giraffe.GetString("db_user"),
		Password: Giraffe.GetString("db_pass"),
	}
	applyUri := Giraffe.GetString("db_host") + ":" + Giraffe.GetString("db_port")

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

type Room interface {
	SetRoomNo(int)
	CheckRoom() (structs.RoomInfo, error)
	GetRoomContents() ([]structs.RoomContents, error)
	CreateRoom(ctx context.Context) error
	SendMessage(ctx context.Context) error
	DeleteRoom(ctx context.Context) error
	DeleteMessage(ctx context.Context) error
	JoinTest(roomNo int)
}

type User interface {
	GetUser(ctx context.Context) (structs.User, error)
}
