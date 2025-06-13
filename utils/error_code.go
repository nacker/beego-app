package utils

const (
	CodeSuccess       = 200
	CodeTokenExpired  = 401
	CodeUserNotFound  = 4000
	CodeBadRequest    = 4001
	CodeUserExisted   = 4002
	CodeUnauthorized  = 4010
	CodeForbidden     = 4030
	CodeNotFound      = 4040
	CodeInternalError = 5000
	CodeDBError       = 5001
	CodeCacheError    = 5002

	// 用户认证错误 (4100-4199)
	CodeInvalidCredentials    = 4100 // 凭据无效(用户名/密码错误)
	CodeAccountDisabled       = 4101 // 账号已禁用
	CodeAccountLocked         = 4102 // 账号已锁定
	CodePasswordExpired       = 4103 // 密码已过期
	CodeTooManyFailedAttempts = 4104 // 登录尝试次数过多
	CodeInvalidCaptcha        = 4105 // 验证码无效
	CodeSessionExpired        = 4106 // 会话已过期
	CodeTwoFactorRequired     = 4107 // 需要二次验证
	CodeTwoFactorInvalid      = 4108 // 二次验证失败
)

var codeMsgMap = map[int]string{
	CodeSuccess:       "success",
	CodeTokenExpired:  "账号已过期，请重新登录",
	CodeUserNotFound:  "用户不存在",
	CodeBadRequest:    "参数错误",
	CodeUserExisted:   "用户已存在",
	CodeUnauthorized:  "未授权",
	CodeForbidden:     "禁止访问",
	CodeNotFound:      "资源不存在",
	CodeInternalError: "服务器内部错误",
	CodeDBError:       "数据库错误",
	CodeCacheError:    "缓存错误",

	CodeInvalidCredentials:    "用户名或密码不正确",
	CodeAccountDisabled:       "账号已被禁用，请联系管理员",
	CodeAccountLocked:         "账号已被锁定，请稍后再试",
	CodePasswordExpired:       "密码已过期，请修改密码",
	CodeTooManyFailedAttempts: "登录尝试次数过多，请稍后再试",
	CodeInvalidCaptcha:        "验证码无效或已过期",
	CodeSessionExpired:        "会话已过期，请重新登录",
	CodeTwoFactorRequired:     "需要进行二次验证",
	CodeTwoFactorInvalid:      "二次验证失败",
}

func MsgForCode(code int) string {
	if msg, ok := codeMsgMap[code]; ok {
		return msg
	}
	return "未知错误"
}
