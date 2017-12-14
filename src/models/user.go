package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"time"
)

type User struct {
	Id       		int64		`orm:"pk;auto"`
	Username 		string		`orm:"size(100);unique"`
	Password 		string		`orm:"size(100)"`
	CreateTime		time.Time	`orm:"auto_now_add;type(datetime)"`
	ModifiedTime	time.Time	`orm:"auto_now_add;type(datetime)"`
}

func (u *User) CreateUser(username, password string) {
	o := orm.NewOrm()
	o.Using("configportal")
	usr := &User{Username:username, Password:password, CreateTime:time.Now(), ModifiedTime:time.Now()}
	o.Insert(usr)
	logs.Info("UserModel CreateUser Success: userName=%v", username)

}

func (u *User) DeleteUser(username string) {
	o := orm.NewOrm()
	o.Using("configportal")
	usr := &User{Username:username}
	o.Delete(usr,"username")
	logs.Info("UserModel DeleteUser Success: userName=%v", usr.Username)

}

func (u *User) UpdateUser(username, password string) {
	o := orm.NewOrm()
	o.Using("configportal")
	usr := &User{Username:username, Password:password, ModifiedTime: time.Now()}
	o.QueryTable("user").Filter("username",username).Update(orm.Params{"password": password, "modified_time": time.Now()})
	logs.Info("UserModel UpdateUser Success: userName=%v", usr.Username)

}

func (u *User) QueryUserByNameAndPassword(name, password string) error {
	// 实例化数据库
	o := orm.NewOrm()
	// 选择要操作的库
	o.Using("configportal")
	user := &User{Username: name, Password: password}
	err := o.Read(user,"username","password")
	*u = *user
	if err != nil {
		logs.Error("UserModel QueryUserByNameAndPassword Error: err = %v", err)
	}
	logs.Info("UserModel QueryUserByNameAndPassword Success: userName = %v,userPass = %v",u.Username, u.Password)
	return err
}
