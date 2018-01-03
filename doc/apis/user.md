用户操作
----------------

# 用户注册

## API
用户注册时请求： POST /api/v1/user

## 请求时提交的数据
请求时提交的请求体如下：

type CreateUserOptions struct {
	UserName	string	`json:"userName"`
	Password	string	`json:"password"`
}

## 返回的数据
返回的数据结构体如下：

type answerData struct {
	Code	int32
	Message	string
	Data	string
}

成功注册时的数据：
Code = OK
Data = ""

前端取数据的key:
result

# 用户注销

## API
用户注销时请求： DELETE /api/v1/user

## 请求时提交的数据
请求时提交的请求体如下：

type DeleteUserOptions struct {
	UserName	string	`json:"userName"`
}

## 返回的数据
返回的数据结构体如下：

type answerData struct {
	Code	int32
	Message	string
	Data	string
}

成功注销时的数据：
Code = OK
Data = ""

前端取数据的key:
result

# 用户修改密码
用户修改密码时请求： PUT /api/v1/user

## 请求时提交的数据
请求时提交的请求体如下：

type UpdateUserOptions struct {
	UserName	string	`json:"userName"`
	Password	string	`json:"password"`
}

## 返回的数据
返回的数据结构体如下：

type answerData struct {
	Code	int32
	Message	string
	Data	string
}

成功修改时的数据：
Code = OK
Data = ""

前端取数据的key:
result
