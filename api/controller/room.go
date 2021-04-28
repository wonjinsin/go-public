package controller

import (
	"context"
	"fmt"
	"gorilla/config"
	"log"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type httpRoomHandler struct {
	db   *mongo.Client
	room *mongo.Collection
}

type JsonInterface struct {
	Test string `json:"test"`
}

func newHTTPRoomHandler(gorilla *config.ViperConfig, eg *echo.Group, db *mongo.Client) {
	h := &httpRoomHandler{
		db:   db,
		room: db.Database("gorilla").Collection("room"),
	}

	eg.GET("", h.Room)
}

func (h *httpRoomHandler) Room(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := h.room.Find(ctx, bson.D{{Key: "roomNumber", Value: "1"}})

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

	return response(c, 200, "test message", "test")
}
