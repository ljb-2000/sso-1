package answerdata

type answerData struct {
	Code	int32
	Message	string
	Data	string
}

func NewAnswer(code int32, data string) *answerData {
	return &answerData {
		Code:	code,
		Message: Messages[code],
		Data:	data,
	}
}

var Messages = map[int32] string {
	OK: "操作成功",
}

const (
	OK int32 = 0
)
