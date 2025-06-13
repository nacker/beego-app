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

func Login(username, password string) (*models.User, error) {
	var user models.User
	err := models.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Register(username string, password string, email string, phone string) (*models.User, error) {

	var user models.User
	user.Username = username
	user.Password = password
	user.Email = email
	user.Phone = phone

	err := models.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
