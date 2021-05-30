package handler

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"giraffe/structs"
	"giraffe/utils"
	"io/ioutil"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type TokenInfo struct {
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
}

type TokenFormat struct {
	*jwt.StandardClaims
	User structs.User
}

func NewTokenHandler() (*TokenInfo, error) {
	signBytes, err := ioutil.ReadFile(fmt.Sprintf("./%s", utils.PrivateKeyPath))
	tokenInfo := &TokenInfo{}

	if err != nil {
		Logger.Logging().Warnw("Fail to Read PrivateKeyFile", "result", err)
		return tokenInfo, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		Logger.Logging().Warnw("Fail to Parse PrivateKeyFile", "result", err)
		return tokenInfo, err
	}

	verifyBytes, err := ioutil.ReadFile(fmt.Sprintf("./%s", utils.PublicKeyPath))

	if err != nil {
		Logger.Logging().Warnw("Fail to Read PublickeyFile", "result", err)
		return tokenInfo, err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		Logger.Logging().Warnw("Fail to Read PublicKeyFile", "result", err)
		return tokenInfo, err
	}

	tokenInfo = &TokenInfo{
		signKey:   signKey,
		verifyKey: verifyKey,
	}

	return tokenInfo, err

}

func (tk *TokenInfo) CreateToken(user structs.User) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims = &TokenFormat{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		user,
	}

	return t.SignedString(tk.signKey)
}

func (tk *TokenInfo) ValidateToken(c echo.Context) (structs.User, error) {
	tokenStr, err := tk.getToken(c)

	if err != nil {
		Logger.Logging().Warnw("Fail to get Baerer token", "result", err)
		return structs.User{}, err
	}

	token, err := jwt.ParseWithClaims(tokenStr, &TokenFormat{}, func(token *jwt.Token) (interface{}, error) {
		return tk.verifyKey, nil
	})

	if err != nil {
		Logger.Logging().Warnw("Fail to parse token", "result", err)
		return structs.User{}, err
	}

	result := token.Claims.(*TokenFormat).User

	return result, nil
}

func (tk *TokenInfo) getToken(c echo.Context) (string, error) {
	bearerToken := c.Request().Header.Get("Authorization")

	strArr := strings.Split(bearerToken, "Bearer ")
	if len(strArr) == 2 {
		return strArr[1], nil
	}

	Logger.Logging().Warnw("Fail to get Baerer token", "result")
	return "", errors.New("no result")
}
