package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id         int       `orm:"auto"`
	Username   string    `orm:"size(50);null"` // 用户名，允许为空
	Email      string    `orm:"size(100);null"` // 电子邮箱，允许为空
	Password   string    `orm:"size(100);null"` // 加密后的密码，允许为空
	Phone      string    `orm:"size(20);null"` // 手机号码，允许为空
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
	UpdateTime time.Time `orm:"auto_now;type(datetime)"` // 更新时间
}

func init() {
	orm.RegisterModel(new(User))
}
