package controllers

import (
	"encoding/json"

	"models"
	"common/gokits/encrypt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"common/gokits/answerdata"
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
	user := cuc.CreateUser(options)
	if user == nil {
		logs.Error("CreateUserController Create User Success: userName=%v", user.Username)
	} else {
		data, _ := json.Marshal(answerdata.NewAnswer(answerdata.OK, ""))
		cuc.Data["result"] = data
	}
}

func (cuc *UserController) CreateUser(options *CreateUserOptions) *models.User{
	user := new(models.User)
	encryptPass := encrypt.NewEncryption(options.Password).String()
	result := user.CreateUser(options.UserName, encryptPass)
	return result
}

func (cuc *UserController) getCreateUserOptionsFromBody() *CreateUserOptions{
	options := new(CreateUserOptions)
	json.Unmarshal(cuc.Ctx.Input.RequestBody, &options)
	logs.Info("UserController Get Body Success", options.UserName, options.Password)
	return options
}


// 删除用户
type DeleteUserOptions struct {
	UserName	string	`json:"userName"`
}

func (cuc *UserController) Delete() {
	options := new(DeleteUserOptions)
	options = cuc.getDeleteUserOptionsFromBody()
	user := cuc.DeleteUser(options)
	if user == nil {
		logs.Error("DeleteUserController Delete User Success: userName=%v", user.Username)
	} else {
		data, _ := json.Marshal(answerdata.NewAnswer(answerdata.OK, ""))
		cuc.Data["result"] = data
	}
}

func (cuc *UserController) getDeleteUserOptionsFromBody() *DeleteUserOptions{
	options := new(DeleteUserOptions)
	json.Unmarshal(cuc.Ctx.Input.RequestBody, &options)
	logs.Info("UserController Get Body Success", options.UserName)
	return options
}

func (cuc *UserController) DeleteUser(options *DeleteUserOptions) *models.User {
	user := new(models.User)
	result := user.DeleteUser(options.UserName)
	return result
}

// 更新用户
type UpdateUserOptions struct {
	UserName	string	`json:"userName"`
	Password	string	`json:"password"`
}

func (cuc *UserController) Put() {
	options := new(UpdateUserOptions)
	options = cuc.getUpdateUserOptionsFromBody()
	user := cuc.UpdateUser(options)
	if user == nil {
		logs.Error("UpdateUserController Update User Success: userName=%v", user.Username)
	} else {
		data, _ := json.Marshal(answerdata.NewAnswer(answerdata.OK, ""))
		cuc.Data["result"] = data
	}
}

func (cuc *UserController) getUpdateUserOptionsFromBody() *UpdateUserOptions{
	options := new(UpdateUserOptions)
	json.Unmarshal(cuc.Ctx.Input.RequestBody, &options)
	logs.Info("UserController Get Body Success: userName=%v", options.UserName)
	return options
}

func (cuc *UserController) UpdateUser(options *UpdateUserOptions) *models.User {
	user := new(models.User)
	encryptPass := encrypt.NewEncryption(options.Password).String()
	result := user.UpdateUser(options.UserName, encryptPass)
	return result
}

