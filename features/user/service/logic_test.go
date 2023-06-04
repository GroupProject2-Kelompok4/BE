package service

import (
	"errors"
	"testing"

	"github.com/GroupProject2-Kelompok4/BE/features/user"
	"github.com/GroupProject2-Kelompok4/BE/mocks"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	data := mocks.NewUserData(t)
	arguments := user.UserCore{Email: "admin@gmail.com", Password: "@SecretPassword123"}
	wrongArguments := user.UserCore{Email: "admin@gmail.com", Password: "@WrongPassword"}
	token := "123"
	emptyToken := ""
	hashed, _ := helper.HashPassword(arguments.Password)
	result := user.UserCore{UserID: "uuid", Fullname: "admin", Password: hashed}
	service := New(data)

	t.Run("email and password cannot be empty", func(t *testing.T) {
		request := user.UserCore{
			Email:    "",
			Password: "",
		}

		_, _, err := service.Login(request)
		expectedErr := errors.New("email and password cannot be empty")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("success login", func(t *testing.T) {
		data.On("Login", mock.Anything).Return(result, token, nil).Once()
		res, token, err := service.Login(arguments)
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, result.Email, res.Email)
		assert.Equal(t, result.Password, res.Password)
		data.AssertExpectations(t)
	})

	t.Run("invalid email and password", func(t *testing.T) {
		data.On("Login", mock.Anything).Return(user.UserCore{}, token, errors.New("invalid email and password")).Once()
		_, _, err := service.Login(wrongArguments)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "invalid email and password")
		data.AssertExpectations(t)
	})

	t.Run("password does not match", func(t *testing.T) {
		data.On("Login", mock.Anything).Return(user.UserCore{}, token, errors.New("password does not match")).Once()
		_, _, err := service.Login(wrongArguments)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password does not match")
		data.AssertExpectations(t)
	})

	t.Run("error while creating jwt token", func(t *testing.T) {
		data.On("Login", mock.Anything).Return(user.UserCore{}, token, errors.New("error while creating jwt token")).Once()
		_, _, err := service.Login(wrongArguments)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error while creating jwt token")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Login", mock.Anything).Return(user.UserCore{}, emptyToken, errors.New("server error")).Once()
		res, token, err := service.Login(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.UserID)
		assert.Equal(t, emptyToken, token)
		assert.ErrorContains(t, err, "internal server error")
		data.AssertExpectations(t)
	})
}
