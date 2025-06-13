package controllers

import (
	"beego-app/services"
	"beego-app/utils"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type UserController struct {
	web.Controller
}

func (c *UserController) GetUser() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	user, err := services.GetUserByID(id)
	if err != nil {
		utils.ErrorCode(c.Ctx, 4040)
		return
	}
	utils.Success(c.Ctx, user)
}
