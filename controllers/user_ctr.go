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
