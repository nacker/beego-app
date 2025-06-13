package controllers

import (
	"beego-app/services"
	"beego-app/utils"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

// UserController 定义用户相关的控制器
type UserController struct {
	web.Controller
}

func (c *UserController) GetUser() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	user, err := services.GetUserByID(id)
	if err != nil {
		utils.Error(c.Ctx, utils.CodeUnauthorized, err.Error())
		return
	}
	utils.Success(c.Ctx, user)
}

func (c *UserController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")
	user, err := services.Login(username, password)
	if err != nil {
		utils.Error(c.Ctx, utils.CodeUnauthorized, err.Error())
		return
	}

	userID := user.ID // 获取用户 ID
	token, err := utils.NewJWT().GenerateToken(strconv.Itoa(userID), user.Username)

	result := map[string]interface{}{
		"token": token,
		"user":  user,
	}
	utils.Success(c.Ctx, result)
}

func (c *UserController) Register() {
	username := c.GetString("username")
	password := c.GetString("password")
	email := c.GetString("email")
	phone := c.GetString("phone")

	user, err := services.Register(username, password, email, phone)
	if err != nil {
		utils.Error(c.Ctx, utils.CodeUnauthorized, err.Error())
		return // ⚠️ 必须 return，防止空指针访问 user
	}

	userID := user.ID // ✅ 确保 User 模型字段名为 ID（首字母大写）
	token, err := utils.NewJWT().GenerateToken(strconv.Itoa(userID), user.Username)
	if err != nil {
		utils.Error(c.Ctx, utils.CodeInternalError, "生成 token 失败: "+err.Error())
		return
	}

	result := map[string]interface{}{
		"token": token,
		"user":  user,
	}
	utils.Success(c.Ctx, result)
}

func (c *UserController) Logout() {
	utils.Success(c.Ctx, nil)
}
