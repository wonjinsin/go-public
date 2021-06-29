package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"giraffe/structs"
	"giraffe/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	e, _ := SetEchoEnv()
	ts = httptest.NewServer(e)

	os.Exit(m.Run())
}

func TestRoom(t *testing.T) {
	assert := assert.New(t)

	resp, err := http.Get(ts.URL + "/api/v1/room/13")

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/api/v1/room/noNum")

	if !assert.NoError(err) {
		utils.PrintError(err)
	}

	assert.Equal(http.StatusNotFound, resp.StatusCode)
}

type CreateRoomStruct struct {
	Room_no string
	User1   string
	User2   string
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	room := CreateRoomStruct{Room_no: "14", User1: "gorilla", User2: "giraffe"}
	pbytes, _ := json.Marshal(room)
	buff := bytes.NewBuffer(pbytes)

	resp, err := http.Post(ts.URL+"/api/v1/room/create", "application/json", buff)
	if !assert.NoError(err) {
		utils.PrintError(err)
	}
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestSend(t *testing.T) {
	assert := assert.New(t)

	roomSend := structs.RoomSend{RoomNo: "13", User: "gorilla", Message: "This is test message"}
	pbytes, _ := json.Marshal(roomSend)
	buff := bytes.NewBuffer(pbytes)

	resp, err := http.Post(ts.URL+"/api/v1/room/send", "application/json", buff)
	if !assert.NoError(err) {
		utils.PrintError(err)
	}
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestDeleteRoom(t *testing.T) {
	assert := assert.New(t)
	roomNo := "14"

	req, err := http.NewRequest("DELETE", ts.URL+fmt.Sprintf("/api/v1/room/delete/%s", roomNo), nil)

	if !assert.NoError(err) {
		utils.PrintError(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestDeleteMessage(t *testing.T) {
	assert := assert.New(t)
	objectId := "60b3a14445f2f7d32a2a33c9"

	req, err := http.NewRequest("DELETE", ts.URL+fmt.Sprintf("/api/v1/room/message/%s", objectId), nil)

	if !assert.NoError(err) {
		utils.PrintError(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if !assert.NoError(err) {
		utils.PrintError(err)
	}
	assert.Equal(http.StatusMethodNotAllowed, resp.StatusCode)
}
