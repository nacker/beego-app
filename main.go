package main

import (
	"beego-app/models"
	_ "beego-app/routers"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
)

func init() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in init:", r)
		}
	}()
	// 路由注册代码...
}

func main() {
	if err := models.InitDB(); err != nil { // 初始化数据库并处理错误
		log.Fatal("Database connection failed: ", err)
	}
	beego.Run()
}
