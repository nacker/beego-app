package models

import (
	"time"
)

type User struct {
	ID         int       `orm:"column(id);auto;pk"`          // 更标准的命名
	Username   string    `orm:"size(50);null"`               // 用户名，允许为空
	Email      string    `orm:"size(100);null"`              // 电子邮箱，允许为空
	Password   string    `orm:"size(100);null"`              // 加密后的密码，允许为空
	Phone      string    `orm:"size(20);null"`               // 手机号码，允许为空
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`     // 更新时间
}

//func init() {
//	// 在数据库连接初始化后执行自动迁移
//	go func() {
//		// 等待数据库连接初始化完成
//		time.Sleep(time.Second)
//		if DB != nil {
//			err := DB.AutoMigrate(&User{})
//			if err != nil {
//				panic("failed to auto migrate user table: " + err.Error())
//			}
//		}
//	}()
//}
