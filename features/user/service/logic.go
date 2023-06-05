package service

import (
	"errors"
	"strings"

	"github.com/GroupProject2-Kelompok4/BE/features/user"
	"github.com/GroupProject2-Kelompok4/BE/utils"
)

var log = utils.Log()

type userService struct {
	query user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userService{
		query: ud,
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

// Register implements user.UserService
func (us *userService) Register(request user.UserCore) (user.UserCore, error) {
	if request.Fullname == "" || request.Email == "" || request.Password == "" {
		log.Error("request cannot be empty")
		return user.UserCore{}, errors.New("request cannot be empty")
	}

	result, err := us.query.Register(request)
	if err != nil {
		message := ""
		if strings.Contains(err.Error(), "error while hashing password") {
			log.Error("error while hashing password")
			message = "error while hashing password"
		} else if strings.Contains(err.Error(), "error insert data, duplicated") {
			log.Error("error insert data, duplicated")
			message = "error insert data, duplicated"
		} else {
			log.Error("internal server error")
			message = "internal server error"
		}
		return user.UserCore{}, errors.New(message)
	}

	return result, nil
}

// SearchUser implements user.UserService
func (us *userService) SearchUser(keyword string, limit, offset int) ([]user.UserCore, uint, error) {
	result, count, err := us.query.SearchUser(keyword, limit, offset)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("not found, error while retrieving list users")
			return []user.UserCore{}, 0, errors.New("not found, error while retrieving list users")
		} else {
			log.Error("internal server error")
			return []user.UserCore{}, 0, errors.New("internal server error")
		}
	}
	return result, count, nil
}
