package controllers

import (
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/database"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/models"
)

func CreateUser(userData models.User) (models.User, error) {
	newUser := userData

	result := database.DB.Create(&newUser)
	if result.Error != nil {
		return newUser, result.Error
	}

	return newUser, nil
}

func GetUserByID(userID string) (models.User, error) {
	var user models.User

	result := database.DB.First(&user, userID)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func UpdateUserByID(userID string, userData models.User) (models.User, error) {
	user, err := GetUserByID(userID)
	if err != nil {
		return user, err
	}

	if userData.Username != "" {
		user.Username = userData.Username
	}

	if userData.Email != "" {
		user.Email = userData.Email
	}

	if userData.Password != "" {
		user.Password = userData.Password
	}

	result := database.DB.Save(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func DeleteUserByID(userID string) error {
	user, err := GetUserByID(userID)
	if err != nil {
		return err
	}

	result := database.DB.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
