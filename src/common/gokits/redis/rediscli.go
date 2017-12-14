package redis

import (
	"github.com/astaxie/beego"
)

func ConnSessionProvider(){
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
}