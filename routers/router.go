package routers

import (
	"beego-app/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/hello", &controllers.MainController{}, "get:Hello")

	beego.Router("/user/:id", &controllers.UserController{}, "get:GetUser")
}
