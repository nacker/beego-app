package utils

const (
	CodeSuccess       = 0
	CodeBadRequest    = 4001
	CodeUnauthorized  = 4010
	CodeNotFound      = 4040
	CodeInternalError = 5000
)

var codeMsgMap = map[int]string{
	CodeSuccess:       "success",
	CodeBadRequest:    "参数错误",
	CodeUnauthorized:  "未授权",
	CodeNotFound:      "资源不存在",
	CodeInternalError: "服务器错误",
}

func MsgForCode(code int) string {
	if msg, ok := codeMsgMap[code]; ok {
		return msg
	}
	return "未知错误"
}
