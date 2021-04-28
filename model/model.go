package model

import (
	"context"
	"fmt"
	"gorilla/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
		log.Fatal(err)
	}

	// Check the connection
	err = db.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Made")
	return db
}
