package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(100)"`
	Age  int
}

func init() {
	orm.RegisterModel(new(User))
}
