package controllers

import (
	"beego-app/utils"
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Index() {
	user := map[string]interface{}{
		"id":   1,
		"name": "Alice",
	}
	utils.Success(c.Ctx, user)
}

func (c *MainController) Hello() {
	utils.Error(c.Ctx, utils.CodeUnauthorized, nil)
}
