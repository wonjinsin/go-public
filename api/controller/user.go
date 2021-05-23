package controller

import (
	"fmt"
	"gorilla/config"
	"gorilla/handler"
	"gorilla/structs"

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

	eg.POST("/login", h.User)
}

func (h *httpUserController) User(c echo.Context) error {
	login := &structs.Login{}
	err := c.Bind(login)

	if err != nil {
		return response(c, 404, "Format is invalid", "")
	}

	fmt.Println(login.Username)

	return response(c, 404, "Not registered yet", "")
}
