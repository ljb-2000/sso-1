package routers

import (
	"controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 用户登录
	beego.Router("/api/v1/login",&controllers.LoginController{})
	// 用户登出
	beego.Router("/api/v1/logout",&controllers.LogoutController{})
	// 验证用户信息
	beego.Router("/api/v1/validation",&controllers.ValidationController{})
	// 用户操作
	beego.Router("/api/v1/user",&controllers.UserController{})
}
