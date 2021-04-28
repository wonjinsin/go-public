package controller

import (
	"gorilla/config"
	"strconv"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type GorillaStatus struct {
	ResultCode string      `json:"resultCode"`
	ResultMsg  string      `json:"resultMsg"`
	ResultData interface{} `json:"resultData,omitempty"`
}

func InitHandler(Gorilla *config.ViperConfig, e *echo.Echo, db *mongo.Client) {

	api := e.Group("/api")
	ver1 := api.Group("/v1")

	room := ver1.Group("/room")
	newHTTPRoomHandler(Gorilla, room, db)
}

func response(c echo.Context, code int, resultMsg string, resultData interface{}) error {

	strCode := strconv.Itoa(code)

	res := GorillaStatus{
		ResultCode: strCode,
		ResultMsg:  resultMsg,
		ResultData: resultData,
	}

	return c.JSON(code, res)

}
