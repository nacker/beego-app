package routers

import (
	"beego-app/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	//beego.InsertFilter("*", beego.BeforeRouter, middlewares.JWTMiddleware)

	//uc := &controllers.UserController{}
	//beego.NSRouter("/", &controllers.MainController{}),
	//beego.NSRouter("/hello", &controllers.MainController{}, "get:Hello"),
	//beego.NSRouter("/user/login", uc, "post:Login"),
	//beego.NSRouter("/user/register", uc, "post:Register"),

	//
	//// 定义用户相关路由命名空间，并应用 JWTMiddleware
	//ns := beego.NewNamespace("/v1",
	//beego.NSBefore(middlewares.JWTMiddleware), // 为该命名空间应用过滤器
	//beego.NSRouter("/:id", &controllers.UserController{}, "get:GetUser"),
	//beego.NSRouter("/logout", &controllers.UserController{}, "post:Logout"),
	//)
	//
	//// 注册命名空间
	//beego.AddNamespace(public)
	//beego.AddNamespace(user)

	// 初始化 namespace
	uc := &controllers.UserController{}
	ns := beego.NewNamespace("/v1",
		// 嵌套 namespace
		beego.NSNamespace("/user",
			beego.NSRouter("/:id", uc, "get:GetUser"),
			beego.NSRouter("/logout", uc, "post:Logout"),
		),
	)
	//注册 namespace
	beego.AddNamespace(ns)
	beego.Run()

}
