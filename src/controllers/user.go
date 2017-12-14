package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"models"
	"common/gokits/encrypt"
)

type UserController struct {
	beego.Controller

}

// 新增用户
type CreateUserOptions struct {
	UserName	string	`json:"userName"`
	Password	string	`json:"password"`
}

func (cuc *UserController) Post() {
	options := new(CreateUserOptions)
	options = cuc.getCreateUserOptionsFromBody()
	cuc.CreateUser(options)

}

func (cuc *UserController) getCreateUserOptionsFromBody() *CreateUserOptions{
	options := new(CreateUserOptions)
	json.Unmarshal(cuc.Ctx.Input.RequestBody, &options)
	logs.Info("UserController Get Body Success", options.UserName, options.Password)
	return options
}

func (cuc *UserController) CreateUser(options *CreateUserOptions) {
	user := new(models.User)
	encryptPass := encrypt.NewEncryption(options.Password).String()
	user.CreateUser(options.UserName, encryptPass)
}

// 删除用户
type DeleteUserOptions struct {
	UserName	string	`json:"userName"`
}

func (cuc *UserController) Delete() {
	options := new(DeleteUserOptions)
	options = cuc.getDeleteUserOptionsFromBody()
	cuc.DeleteUser(options)

}

func (cuc *UserController) getDeleteUserOptionsFromBody() *DeleteUserOptions{
	options := new(DeleteUserOptions)
	json.Unmarshal(cuc.Ctx.Input.RequestBody, &options)
	logs.Info("UserController Get Body Success", options.UserName)
	return options
}

func (cuc *UserController) DeleteUser(options *DeleteUserOptions) {
	user := new(models.User)
	user.DeleteUser(options.UserName)
}

// 更新用户
type UpdateUserOptions struct {
	UserName	string	`json:"userName"`
	Password	string	`json:"password"`
}

func (cuc *UserController) Put() {
	options := new(UpdateUserOptions)
	options = cuc.getUpdateUserOptionsFromBody()
	cuc.UpdateUser(options)

}

func (cuc *UserController) getUpdateUserOptionsFromBody() *UpdateUserOptions{
	options := new(UpdateUserOptions)
	json.Unmarshal(cuc.Ctx.Input.RequestBody, &options)
	logs.Info("UserController Get Body Success: userName=%v", options.UserName)
	return options
}

func (cuc *UserController) UpdateUser(options *UpdateUserOptions) {
	user := new(models.User)
	encryptPass := encrypt.NewEncryption(options.Password).String()
	user.UpdateUser(options.UserName, encryptPass)
}

