package middlewares

import (
	"github.com/beego/beego/v2/server/web/context"
)

func JWTMiddleware(ctx *context.Context) {
	//token := ctx.Input.Header("Authorization")
	//if token != "Bearer expected_token" {
	//	ctx.Output.SetStatus(401)
	//	ctx.Output.Body([]byte("Unauthorized"))
	//}
}
