package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoom(t *testing.T) {
	assert := assert.New(t)

	e, _ := SetEchoEnv()
	ts := httptest.NewServer(e)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/v1/room/13")

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/api/v1/room/noNum")

	assert.NoError(err)
	assert.Equal(http.StatusNotFound, resp.StatusCode)

}
