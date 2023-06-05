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

func TestRegister(t *testing.T) {
	data := mocks.NewUserData(t)
	arguments := user.UserCore{
		Fullname: "admin",
		Email:    "admin@gmail.com",
		Password: "@S3#cr3tP4ss#word123",
		Status:   "mentor",
		Role:     "user",
	}
	result := user.UserCore{
		UserID:   "550e8400-e29b-41d4-a716-446655440000",
		Fullname: "admin",
		Email:    "admin@gmail.com",
		Password: "@S3#cr3tP4ss#word123",
		Status:   "mentor",
		Role:     "user",
	}
	service := New(data)

	t.Run("request cannot be empty", func(t *testing.T) {
		request := user.UserCore{
			Fullname: "admin",
			Email:    "081235288543",
			Password: "",
			Status:   "",
		}
		_, err := service.Register(request)
		expectedErr := errors.New("request cannot be empty")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("success create account", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(result, nil).Once()
		res, err := service.Register(arguments)
		assert.Nil(t, err)
		assert.Equal(t, result.UserID, res.UserID)
		assert.NotEmpty(t, result.Fullname)
		assert.NotEmpty(t, result.Email)
		assert.NotEmpty(t, result.Password)
		data.AssertExpectations(t)
	})

	t.Run("error while hashing password", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.UserCore{}, errors.New("error while hashing password")).Once()
		res, err := service.Register(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.UserID)
		assert.ErrorContains(t, err, "password")
		data.AssertExpectations(t)
	})

	t.Run("error insert data, duplicated", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.UserCore{}, errors.New("error insert data, duplicated")).Once()
		res, err := service.Register(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.UserID)
		assert.ErrorContains(t, err, "duplicated")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.UserCore{}, errors.New("server error")).Once()
		res, err := service.Register(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.UserID)
		assert.ErrorContains(t, err, "internal server error")
		data.AssertExpectations(t)
	})
}

func TestSearchUser(t *testing.T) {
	data := mocks.NewUserData(t)
	service := New(data)
	keyword := "admin"
	limit := 10
	offset := 0
	expectedResult := []user.UserCore{
		{UserID: "550e8400-e29b-41d4-a716-446655440000", Fullname: "admin", Email: "admin@mail.com"},
	}
	expectedCount := uint(1)

	t.Run("success", func(t *testing.T) {

		data.On("SearchUser", keyword, limit, offset).Return(expectedResult, expectedCount, nil)

		result, _, err := service.SearchUser(keyword, limit, offset)

		assert.Nil(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, expectedResult[0].UserID, result[0].UserID)
		assert.Equal(t, expectedResult[0].Fullname, result[0].Fullname)
		assert.Equal(t, expectedResult[0].Email, result[0].Email)
		data.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		data.On("SearchUser", keyword, limit, offset).Return([]user.UserCore{}, uint(0), errors.New("not found, error while retrieving list users")).Once()

		result, _, err := service.SearchUser(keyword, limit, offset)

		assert.NotNil(t, uint(0), err)
		assert.Empty(t, result)
		assert.EqualError(t, err, "not found, error while retrieving list users")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("SearchUser", keyword, limit, offset).Return([]user.UserCore{}, uint(0), errors.New("internal server error")).Once()

		result, _, err := service.SearchUser(keyword, limit, offset)

		assert.NotNil(t, uint(0), err)
		assert.Empty(t, result)
		assert.EqualError(t, err, "internal server error")
		data.AssertExpectations(t)
	})
}

func TestProfile(t *testing.T) {
	data := mocks.NewUserData(t)
	result := user.UserCore{
		UserID:   "550e8400-e29b-41d4-a716-446655440000",
		Fullname: "admin",
		Email:    "admin@mail.com"}
	service := New(data)

	t.Run("success get profile", func(t *testing.T) {
		data.On("ProfileUser", "550e8400-e29b-41d4-a716-446655440000").Return(result, nil).Once()
		res, err := service.ProfileUser("550e8400-e29b-41d4-a716-446655440000")
		assert.Nil(t, err)
		assert.Equal(t, result.UserID, res.UserID)
		assert.Equal(t, result.Fullname, res.Fullname)
		assert.Equal(t, result.Email, res.Email)
		data.AssertExpectations(t)
	})

	t.Run("not found, error while retrieving user profile", func(t *testing.T) {
		data.On("ProfileUser", "550e8400-e29b-41d4-a716-446655440000").Return(user.UserCore{}, errors.New("not found, error while retrieving user profile")).Once()
		res, err := service.ProfileUser("550e8400-e29b-41d4-a716-446655440000")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, user.UserCore{}, res)
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("ProfileUser", "550e8400-e29b-41d4-a716-446655440000").Return(user.UserCore{}, errors.New("internal server error")).Once()
		res, err := service.ProfileUser("550e8400-e29b-41d4-a716-446655440000")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		assert.Equal(t, user.UserCore{}, res)
		data.AssertExpectations(t)
	})
}
