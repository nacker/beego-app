package middlewares

import (
	"beego-app/utils"
	"github.com/beego/beego/v2/server/web/context"
	"log"
	"strings"
)

// JWTMiddleware 处理JWT身份验证的中间件
func JWTMiddleware(ctx *context.Context) {

	// 排除不需要验证的路由
	excludePaths := []string{"/", "/hello", "/user/login", "/user/register"}
	for _, path := range excludePaths {
		if ctx.Request.URL.Path == path {
			return // 跳过中间件逻辑
		}
	}

	// 从请求头中获取 Authorization 字段
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		//ctx.Output.SetStatus(401)
		log.Printf("[JWT] 未提供 Token") // 添加日志记录
		utils.ErrorCode(ctx, 401, "未提供 Token", nil)
		return
	}

	// 分割 Authorization 字段，获取 Token
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		log.Printf("[JWT] 无效的 Token 格式: %s", authHeader) // 添加日志记录
		//ctx.Output.SetStatus(401)
		utils.ErrorCode(ctx, 401, "无效的 Token 格式", nil)

		return
	}

	tokenString := authParts[1]

	// 初始化 JWT 实例
	jwtInstance := utils.NewJWT()

	// 解析并验证 Token
	claims, err := jwtInstance.ParseToken(tokenString)
	if err != nil {
		//ctx.Output.SetStatus(401)
		log.Printf("[JWT] 无效的 Token: %v", err) // 添加日志记录
		utils.ErrorCode(ctx, 401, "无效的 Token: ", err.Error())
		return
	}

	// 将用户信息存入上下文
	ctx.Input.SetData("userID", claims.UserID)
	ctx.Input.SetData("userName", claims.UserName)
	log.Printf("[JWT] 验证成功 - 用户ID: %s, 用户名: %s", claims.UserID, claims.UserName) // 添加日志记录

}
