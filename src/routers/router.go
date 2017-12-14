package routers

import (
	"controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/v1/login",&controllers.LoginController{})
	beego.Router("/api/v1/validation",&controllers.ValidationController{})
	beego.Router("/api/v1/user",&controllers.UserController{})
}
