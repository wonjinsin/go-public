package model

import (
	"giraffe/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client
var rm *RoomModel

func TestMain(m *testing.M) {
	Giraffe := config.Giraffe

	db = MongoConn(Giraffe)
	rm = NewRoomModel(db)
	os.Exit(m.Run())
}

func TestCheckRoom(t *testing.T) {
	assert := assert.New(t)

	rm.SetRoomNo(13)
	_, err := rm.CheckRoom()

	assert.NoError(err)
}
