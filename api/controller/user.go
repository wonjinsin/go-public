package controller

import (
	"context"
	"gorilla/config"
	"gorilla/handler"
	"gorilla/structs"
	"gorilla/utils"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type httpUserController struct {
	db *mongo.Client
	uh handler.User
}

func newHTTPUserContoller(gorilla *config.ViperConfig, eg *echo.Group, db *mongo.Client) {
	uh := handler.NewUserHandler(db)

	h := &httpUserController{
		db: db,
		uh: uh,
	}

	eg.POST("/login", h.Login)
}

func (h *httpUserController) Login(c echo.Context) error {
	login := &structs.Login{}
	err := c.Bind(login)

	if err != nil {
		return response(c, 404, "Format is invalid", "")
	}

	var key utils.StringKey = "loginInfo"

	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, key, login)

	err = h.uh.Login(ctx)

	if err != nil {
		return response(c, 404, "Fail to login", login)
	}

	return response(c, 404, "Not registered yet", "")
}
