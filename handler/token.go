package handler

import (
	"crypto/rsa"
	"fmt"
	"gorilla/structs"
	"gorilla/utils"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenInfo struct {
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
}

type TokenFormat struct {
	*jwt.StandardClaims
	structs.User
}

type TokenUser struct {
	Name string
	Age  string
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

func (tk *TokenInfo) createToken(user structs.User) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	fmt.Println(user)

	t.Claims = &TokenFormat{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		user,
	}

	return t.SignedString(tk.signKey)
}
