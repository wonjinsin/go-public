package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/api/v1/room/noNum")

	assert.NoError(err)
	assert.Equal(http.StatusNotFound, resp.StatusCode)
}

type CreateRoomStruct struct {
	Room_no string
	User1   string
	User2   string
}

func TestCreate(t *testing.T) {
	room := CreateRoomStruct{Room_no: "14", User1: "gorilla", User2: "giraffe"}
	pbytes, _ := json.Marshal(room)
	buff := bytes.NewBuffer(pbytes)

	assert := assert.New(t)

	resp, err := http.Post(ts.URL+"/api/v1/room/create", "application/json", buff)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestDeleteRoom(t *testing.T) {
	roomNo := "14"
	assert := assert.New(t)

	req, err := http.NewRequest("DELETE", ts.URL+fmt.Sprintf("/api/v1/room/delete/%s", roomNo), nil)

	assert.NoError(err)

	client := &http.Client{}
	resp, err := client.Do(req)

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}
