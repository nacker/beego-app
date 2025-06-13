package routers

import (
	"beego-app/controllers"
	"beego-app/middlewares"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.BConfig.RunMode = "dev"
	//beego.BConfig.Listen.EnableAdmin = true
	//beego.BConfig.Listen.AdminAddr = "localhost"
	//beego.BConfig.Listen.AdminPort = 8088

	//beego.InsertFilter("*", beego.BeforeRouter, middlewares.JWTMiddleware)

	// 注册根路径的 GET 请求到 MainController.Index()
	beego.Router("/", &controllers.MainController{}, "get:Index")
	// 注册 /hello 路径的 GET 请求到 MainController.Hello()
	beego.Router("/hello", &controllers.MainController{}, "get:Hello")
	// 注册用户登录和注册路由
	uc := &controllers.UserController{}
	beego.NSRouter("/auth/login", uc, "post:Login")
	beego.NSRouter("/auth/register", uc, "post:Register")

	// 定义用户相关路由命名空间，暂时注释中间件
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(middlewares.JWTMiddleware), // 为该命名空间应用过滤器
		beego.NSRouter("/user/:id", &controllers.UserController{}, "get:GetUser"),
		beego.NSRouter("/user/logout", &controllers.UserController{}, "post:Logout"),
	)

	// 注册命名空间
	beego.AddNamespace(ns)

}
