package models

import (
	"github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dbUser, _ := web.AppConfig.String("dbuser")
	dbPassword, _ := web.AppConfig.String("dbpassword")
	dbHost, _ := web.AppConfig.String("dbhost")
	dbPort, _ := web.AppConfig.String("dbport")
	dbName, _ := web.AppConfig.String("dbname")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return nil
}
