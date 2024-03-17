package helpers

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

type UserReq struct {
	Username string `json:"username" valid:"required~username is required"`
	Email    string `json:"email" valid:"email,required~email is required"`
	Password string `json:"password" valid:"required~password is required"`
}

func ValidateUserData(userData UserReq) error {
	_, err := govalidator.ValidateStruct(userData)
	if err != nil {
		return errors.New("error validating user data")
	}

	if len(userData.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}
