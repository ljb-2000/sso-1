package main

import (
	_ "routers"
	 _ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego"
	"common/gokits/mysql"
	"github.com/astaxie/beego/orm"
	"models"
	"common/gokits/redis"
	"encoding/gob"
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

//	var hmacSampleSecret []byte
//var FilterUser = func(ctx *context.Context) {
	// sample token string taken from the New example
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		// Don't forget to validate the alg is what you expect:
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//		}
//
//		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
//		return hmacSampleSecret, nil
//	})
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		fmt.Println(claims["foo"], claims["nbf"])
//	} else {
//		fmt.Println("err",err)
//	}
//
//
//}

	//var FilterUser = func(ctx *context.Context) {
	//	tokenString := ctx.Input.Header("Authorization")
	//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	//		}
	//		return hmacSampleSecret, nil
	//	})
	//logs.Info("MainController :claims=%v, ok=%v",jwt.MapClaims{})
	//	if claims, ok := token.Claims.(jwt.MapClaims); ok  {
	//		fmt.Println(claims["sessionId"], claims["userId"])
	//	} else {
	//		fmt.Println(err)
	//	}
	//}
	//
	//var FilterUser = func(ctx *context.Context) {
	//	sessionId := ctx.Input.Header("Authorization")
	//	userId := ctx.Input.Header("UserId")
	//	uId, _ := strconv.Atoi(userId)
	//	if ctx.Input.Session(sessionId) == nil && ctx.Request.RequestURI != "/api/v1/login" {
	//		ctx.Redirect(302, "/api/v1/login")
	//	} else if ctx.Input.Session(sessionId) != nil && ctx.Request.RequestURI != "/api/v1/login"{
	//		session := ctx.Input.Session(sessionId).(*models.Session)
	//		if session.UserId != int64(uId) {
	//			ctx.Redirect(302, "/api/v1/login")
	//		}
	//	}
	//}

	//beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)
	beego.Run()
}
