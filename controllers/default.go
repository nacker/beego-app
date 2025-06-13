package controllers

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	web.Controller
}

func (c *MainController) Index() {
	c.Data["Website"] = "https://beego.vip"
	c.TplName = "index.tpl"
}

func (c *MainController) Hello() {
	c.Ctx.WriteString("Hello Beego v2!")
}
