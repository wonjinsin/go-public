package controller

import (
	"giraffe/config"
	"giraffe/utils"
	"strconv"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

var Logger *utils.Logger

type GiraffeStatus struct {
	ResultCode string      `json:"resultCode"`
	ResultMsg  string      `json:"resultMsg"`
	ResultData interface{} `json:"resultData,omitempty"`
}

func InitHandler(Giraffe *config.ViperConfig, e *echo.Echo, db *mongo.Client) {

	api := e.Group("/api")
	ver1 := api.Group("/v1")

	room := ver1.Group("/room")
	user := ver1.Group("/user")
	newHTTPRoomContoller(Giraffe, room, db)
	newHTTPUserContoller(Giraffe, user, db)
}

func response(c echo.Context, code int, resultMsg string, resultData interface{}) error {
	strCode := strconv.Itoa(code)

	res := GiraffeStatus{
		ResultCode: strCode,
		ResultMsg:  resultMsg,
		ResultData: resultData,
	}

	return c.JSON(code, res)
}
