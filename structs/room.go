package structs

import "time"

type RoomInfo struct {
	RoomNo   int            `json:"roomNo"`
	Users    [2]string      `json:"users"`
	Contents []RoomContents `json:"contents"`
}

type RoomContents struct {
	User    string    `json:"user"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
	DateStr string    `json:"DateStr"`
}

type RoomSend struct {
	RoomNo  string `json:"room_no"`
	User    string `json:"user"`
	Message string `json:"message"`
}

type RoomSendInfo struct {
	RoomNo  int       `bson:"roomNo"`
	User    string    `bson:"user"`
	Message string    `bson:"message"`
	Date    time.Time `bson:"date"`
}

type RoomCreate struct {
	RoomNo string `json:"room_no"`
	User1  string `json:"user1"`
	User2  string `json:"user2"`
}

type RoomCreateInfo struct {
	RoomNo int       `bson:"roomNo"`
	Users  [2]string `bson:"users"`
}

type RoomDeleteInfo struct {
	ObjectId string `bson:"objectId"`
}
