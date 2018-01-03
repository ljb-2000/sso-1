用户登录
----------------

# 用户登录

## API
用户登录时请求： POST /api/v1/login

## 请求时提交的数据
请求时提交的请求体如下：

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

## 返回的数据
返回的数据结构体如下：

type answerData struct {
	Code	int32
	Message	string
	Data	string
}

用户成功登录时的数据：
返回token
Code = OK
Data = tokenString

前端取数据的key:
token

返回username
Code = OK
Data = UserName

前端取数据的key:
username

