package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"models"
	"common/gokits/encrypt"
	"github.com/dgrijalva/jwt-go"
	"common/gokits/answerdata"
)

type LoginController struct {
	beego.Controller
}

// 用户登录参数
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 入口函数
func (lc *LoginController) Post() {
	// 初始化用户登录参数
	var loginParams LoginParams
	// 获取用户登录参数
	loginParams = lc.getBodyOptions()
	// 检查数据库中有无用户信息
	user := lc.check(loginParams.Username, loginParams.Password)
	if user == nil {
		logs.Error("LoginController Check Error: userName = %v, userPass = %v",loginParams.Username, loginParams.Password)
	} else {
		logs.Info("LoginController Check Success: userName = %v, userPass = %v",user.Username,user.Password)
	}
	// 建立用户会话
	lc.session(user)

}

func (lc *LoginController) getBodyOptions() LoginParams{
	var options LoginParams
	json.Unmarshal(lc.Ctx.Input.RequestBody, &options)
	logs.Info("LoginController Get Body Success: username = %s, password = %s", options.Username, options.Password)
	return options
}

// 检查数据库中有无用户信息
func (lc *LoginController) check(name, password string) *models.User {
	// 加密用户密码
	encryptPass := encrypt.NewEncryption(password).String()
	// 初始化用户信息
	user := new(models.User)
	// 通过用户名与密码查询用户信息
	err := user.QueryUserByNameAndPassword(name, encryptPass)
	// 查询失败返回nil
	if err != nil {
		logs.Error("LoginController QueryUserByNameAndPassword Error: error=%s, userName=%s", err, name)
		return nil
	}
	logs.Info("LoginController Check Success: userName=%v", user.Username)
	// 返回用户信息
	return user
}

// 建立会话并返回token
func (lc *LoginController) session(user *models.User) {
	// 秘钥
	var secret = "config"
	// 新建会话
	session := new(models.Session)
	session = models.NewSession(user.Id, user.Username)
	// 会话写入redis中
	lc.SetSession(session.SessionId, session)
	// sessionId与userId做成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sessionId": session.SessionId, "userId": session.UserId})
	if token != nil {
		tokenString, _ := token.SignedString([]byte(secret))
		if tokenString != "" {
			// 将token返回给前端
			data, _ := json.Marshal(answerdata.NewAnswer(answerdata.OK, tokenString))
			lc.Data["token"] = data
			// 将username返回给前端
			data, _ = json.Marshal(answerdata.NewAnswer(answerdata.OK, user.Username))
			lc.Data["username"] = data
		}
		logs.Info("LoginController token return Success: tokenString=%v, username=%v", tokenString, user.Username)
	} else {
		logs.Error("LoginController token return Error: sessionId=%v, username=%v", session.SessionId, session.UserName)
	}
}