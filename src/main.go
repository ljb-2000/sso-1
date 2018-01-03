package main

import (
	"fmt"
	"encoding/gob"
	"strconv"

	_ "routers"
	"models"
	"common/gokits/mysql"
	"common/gokits/redis"

	_ "github.com/astaxie/beego/session/redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

func init() {
	gob.Register(new(models.Session))
	mysql.ConnMysql()
	redis.ConnSessionProvider()
	orm.RegisterModel(new(models.User))
	orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 验证、过滤
	var FilterUser = func(ctx *context.Context) {
		var ok bool
		var secret = []byte("config")
		userId := ctx.Input.Header("UserId")
		uId, _ := strconv.Atoi(userId)
		tokenString := ctx.Input.Header("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		 _, ok = token.Method.(*jwt.SigningMethodHMAC)
		 if !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				logs.Info("MainController :claims=%v, ok=%v",jwt.MapClaims{},ok)
			}
			logs.Info("MainController :claims=%v, ok=%v, secret=%v",jwt.MapClaims{}, ok, string(secret))
			return secret, nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok  && token.Valid {
			fmt.Println(claims["sessionId"], claims["userId"])
			sessionId := claims["sessionId"]
		if ctx.Input.Session(sessionId) == nil && ctx.Request.RequestURI != "/api/v1/login" {
			ctx.WriteString("Validation Error: Session nil")
		} else if ctx.Input.Session(sessionId) != nil && ctx.Request.RequestURI != "/api/v1/login"{
			session := ctx.Input.Session(sessionId).(*models.Session)
			if session.UserId != int64(uId) {
				ctx.WriteString("Validation Error: UserId Error")
			}
		}
		} else {
			ctx.WriteString("Validation Error: Token Error")
			fmt.Println(err)
		}
	}

	// 过滤器
	beego.InsertFilter("/api/v1/validation", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/api/v1/logout", beego.BeforeRouter, FilterUser)

	beego.Run()
}
