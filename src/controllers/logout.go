package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"common/gokits/answerdata"
	"encoding/json"
)

type LogoutController struct {
	beego.Controller
}

// 用户登出参数
type LogoutParams struct {
	tokenString string
}

// 入口函数
func (lc *LogoutController) Post() {
	// 初始化用户登出参数
	var logoutParams LogoutParams
	// 请求头中读出token
	logoutParams.tokenString= lc.Ctx.Input.Header("Authorization")
	if logoutParams.tokenString == "" {
		logs.Error("LogoutController Authorization Error: Not find Authorization")
		return
	}
	result := lc.deleteSession(logoutParams.tokenString)
	if result == true{
		data, _ := json.Marshal(answerdata.NewAnswer(answerdata.OK, ""))
		lc.Data["result"] = data
	}
}

// 删除会话
func (lc *LogoutController) deleteSession(tokenString string) bool {
	var ok bool
	// 秘钥
	var secret = []byte("config")
	var sessionId string

	// 解析出token
	token, _:= jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		 _, ok = token.Method.(*jwt.SigningMethodHMAC)
		 if !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				logs.Info("LogoutController Token Parse Success:claims=%v, ok=%v",jwt.MapClaims{},ok)
			}
			logs.Info("LogoutController Token Parse Success:claims=%v, ok=%v, secret=%v",jwt.MapClaims{}, ok, string(secret))
			return secret, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok  && token.Valid {
		fmt.Println(claims["sessionId"], claims["userId"])
		sessionId = claims["sessionId"].(string)
		// 删除会话
		lc.DelSession(sessionId)
		return true
	} else {
			logs.Error("LogoutController Delete Session Error: sessionId=%v", claims["sessionId"])
			return false
	}
}


