用户登出
----------------

# 用户登出

## API
用户登出时请求： POST /api/v1/logout

## 请求时提交的数据
请求时提交的请求头如下：

Authorization
UserId

## 返回的数据
返回的数据结构体如下：

type answerData struct {
	Code	int32
	Message	string
	Data	string
}

用户成功登出时的数据：
Code = OK
Data = ""

前端取数据的key:
result

