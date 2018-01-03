package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

type User struct {
	Id       		int64		`orm:"pk;auto"`
	Username 		string		`orm:"size(100);unique"`
	Password 		string		`orm:"size(100)"`
	CreateTime		time.Time	`orm:"auto_now_add;type(datetime)"`
	ModifiedTime	time.Time	`orm:"auto_now_add;type(datetime)"`
}

// 创建用户
func (u *User) CreateUser(username, password string) *User{
	o := orm.NewOrm()
	o.Using("configportal")
	usr := &User{Username:username, Password:password, CreateTime:time.Now(), ModifiedTime:time.Now()}
	_, err := o.Insert(usr)
	if err != nil{
		logs.Error("UserModel CreateUser Error: userName=%v", username)
		return nil
	}
	logs.Info("UserModel CreateUser Success: userName=%v", username)
	return usr

}

// 删除用户
func (u *User) DeleteUser(username string) *User{
	o := orm.NewOrm()
	o.Using("configportal")
	usr := &User{Username:username}
	_, err := o.Delete(usr,"username")
	if err != nil {
		logs.Error("UserModel CreateUser Error: userName=%v", username)
		return nil
	}
	logs.Info("UserModel DeleteUser Success: userName=%v", usr.Username)
	return usr
}

// 更新用户
func (u *User) UpdateUser(username, password string) *User{
	o := orm.NewOrm()
	o.Using("configportal")
	usr := &User{Username:username, Password:password, ModifiedTime: time.Now()}
	_, err := o.QueryTable("user").Filter("username",username).Update(orm.Params{"password": password, "modified_time": time.Now()})
	if err != nil {
		logs.Error("UserModel CreateUser Error: userName=%v", username)
		return nil
	}
	logs.Info("UserModel UpdateUser Success: userName=%v", usr.Username)
	return usr
}

// 通过用户名与密码查询用户
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
	} else {
		logs.Info("UserModel QueryUserByNameAndPassword Success: userName = %v,userPass = %v", u.Username, u.Password)
	}
	return err
}
