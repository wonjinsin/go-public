package controller

import (
	"context"
	"gorilla/config"
	"gorilla/handler"
	"gorilla/structs"
	"gorilla/utils"
	"net/http"
	"strings"

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
	eg.GET("/validate", h.Validate)
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

	token, err := h.uh.Login(ctx)

	if err != nil {
		return response(c, 404, "Fail to login", login)
	}

	return response(c, 200, "Login Succeed", token)
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (h *httpUserController) Extract(c echo.Context) error {
	return response(c, 200, "Validate success", "")
}

func (h *httpUserController) Validate(c echo.Context) error {
	user, err := h.uh.Validate(c)

	if err != nil {
		return response(c, 400, "Validate Failed", err)
	}

	return response(c, 200, "Validate success", user)
}
