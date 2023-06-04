package service

import (
	"errors"
	"strings"

	"github.com/GroupProject2-Kelompok4/BE/features/user"
	"github.com/GroupProject2-Kelompok4/BE/utils"
	"github.com/go-playground/validator/v10"
)

var log = utils.Log()

type userService struct {
	query    user.UserData
	validate *validator.Validate
}

func New(ud user.UserData) user.UserService {
	return &userService{
		query:    ud,
		validate: validator.New(),
	}
}

// Login implements user.UserService
func (us *userService) Login(request user.UserCore) (user.UserCore, string, error) {
	if request.Email == "" || request.Password == "" {
		log.Error("email and password cannot be empty")
		return user.UserCore{}, "", errors.New("email and password cannot be empty")
	}

	result, token, err := us.query.Login(request)
	if err != nil {
		message := ""
		if strings.Contains(err.Error(), "invalid email and password") {
			log.Error("invalid email and password")
			message = "invalid email and password"
		} else if strings.Contains(err.Error(), "password does not match") {
			log.Error("password does not match")
			message = "password does not match"
		} else if strings.Contains(err.Error(), "error while creating jwt token") {
			log.Error("error while creating jwt token")
			message = "error while creating jwt token"
		} else {
			log.Error("internal server error")
			message = "internal server error"
		}
		return user.UserCore{}, "", errors.New(message)
	}

	return result, token, nil
}
