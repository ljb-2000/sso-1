package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"common/config"
)

func ConnMysql() {
	orm.Debug = true
	// 初始化连接
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 创建连接所需的配置数据
	config := new(config.Config)
	config.Instance()
	conn := config.GetDbEndpoint()
	// 连接数据库
	orm.RegisterDataBase("default", "mysql", conn)
}
