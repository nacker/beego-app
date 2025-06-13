package utils

import (
	"github.com/beego/beego/v2/server/web/context"
)

// 通用响应结构
type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 成功响应
func Success(ctx *context.Context, data interface{}) {
	resp := ApiResponse{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	if err := ctx.Output.JSON(resp, false, false); err != nil {
		// 写日志更合适
	}
}

// 错误响应
func Error(ctx *context.Context, code int, message string) {
	resp := ApiResponse{
		Code:    code,
		Message: message,
	}
	if err := ctx.Output.JSON(resp, false, false); err != nil {
		// 写日志更合适

	}
}

func ErrorCode(ctx *context.Context, code int) {
	Error(ctx, code, MsgForCode(code))
}
