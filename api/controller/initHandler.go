package controller

import (
	"giraffe/config"
	"giraffe/model"
	"giraffe/utils"
	"net/http"
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

func SetEchoEnv() (*echo.Echo, *config.ViperConfig) {
	Giraffe := config.Giraffe
	e := echo.New()
	e.GET("/healthCheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "It's working!")
	})

	db := model.MongoConn(Giraffe)
	initHandler(Giraffe, e, db)

	return e, Giraffe
}

func initHandler(Giraffe *config.ViperConfig, e *echo.Echo, db *mongo.Client) {

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
