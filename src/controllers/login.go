package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"models"
	"common/gokits/encrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type LoginController struct {
	beego.Controller
}

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (lc *LoginController) Post() {
	var loginParams LoginParams
	loginParams = lc.getBodyOptions()
	user := lc.check(loginParams.Username, loginParams.Password)
	logs.Info("LoginController Check Success: userName = %v, userPass = %v",user.Username,user.Password)
	lc.session(user)

}

func (lc *LoginController) getBodyOptions() LoginParams{
	var options LoginParams
	json.Unmarshal(lc.Ctx.Input.RequestBody, &options)
	logs.Info("LoginController Get Body Success: username = %s, password = %s", options.Username, options.Password)
	return options
}

func (lc *LoginController) check(name, password string) *models.User {
	encryptPass := encrypt.NewEncryption(password).String()

	user := new(models.User)
	err := user.QueryUserByNameAndPassword(name, encryptPass)
	if err != nil {
		logs.Error("LoginController QueryUserByNameAndPassword Error: error=%s, userName=%s", err, name)
		return nil
	}
	logs.Info("LoginController Check Success: userName=%v", user.Username)
	return user
}

func (lc *LoginController) session(user *models.User) {
	var hmacSampleSecret []byte
	session := new(models.Session)
	session = models.NewSession(user.Id, user.Username)
	lc.SetSession(session.SessionId, session)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"exp": 3600, "userId": session.UserId, "timeStamp": time.Now().String()})
	tokenString, _ := token.SignedString(hmacSampleSecret)
	logs.Info("LoginController token return Success: tokenString=%v, token=%v", tokenString, token)
	lc.Data["token"] = tokenString
	//logs.Info("LoginController session return Success: value=%v", session)
	//lc.Data["session"] = session
}