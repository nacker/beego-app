package main

import (
	_ "beego-app/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
}

func init() {
	//// 注册模型
	//orm.RegisterModel(new(User))
	//// need to register default database
	//orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:password@tcp(127.0.0.1:3306)/api?charset=utf8")
}

func main() {
	beego.Run()

}
