package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Login struct {
	Username string `json:"name" bson:"name"`
}

type User struct {
	ObjectId primitive.ObjectID `json:"objectId"`
	Age      int                `json:"age"`
	Name     string             `json:"name"`
}
