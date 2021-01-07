package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type AuthController struct {
	beego.Controller
}

func IssueJWT(userName, phonenum string, birthday time.Time) string {
	token := jwt.New(jwt.SigningMethodHS256)

	//JWTの期限はは１ヶ月に設定しておく(もう少し期限の猶予を追加してもいいかも)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userName
	claims["birthdate"] = birthday
	claims["TeDeUserPhoneNum"] = phonenum
	claims["iss"] = beego.AppConfig.String("appname")
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 730).Unix()

	tokenString, _ := token.SignedString([]byte(beego.AppConfig.String("JwtSecretKey")))

	return tokenString
}
