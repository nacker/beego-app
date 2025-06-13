package services

import (
	"beego-app/models"
)

func GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := models.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//func CreateUser(name string, age int) (*models.User, error) {
//	var user models.User
//	err := models.DB.Create(&models.User{Name: name, Age: age}).Error
//	if err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
