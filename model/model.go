package model

import (
	"context"
	"fmt"
	"gorilla/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type logEntity struct {
	logType  string `bson:"logType"`
	logValue string `bson:"logValue"`
	dbPrefix string `bson:"dbPrefix"`
	time     string `bson:"time"`
}

func MongoConn(gorilla *config.ViperConfig) (client *mongo.Client) {

	credential := options.Credential{
		Username: gorilla.GetString("db_user"),
		Password: gorilla.GetString("db_pass"),
	}
	applyUri := gorilla.GetString("db_host") + ":" + gorilla.GetString("db_port")

	clientOptions := options.Client().ApplyURI("mongodb://" + applyUri).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Made")
	return client
}
