package service

import (
	"errors"
	"testing"
	"time"

	"github.com/GroupProject2-Kelompok4/BE/features/class"
	"github.com/GroupProject2-Kelompok4/BE/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterClass(t *testing.T) {
	data := mocks.NewClassData(t)
	arguments := class.ClassCore{
		ClassID:      "550e8400-e29b-41d4-a716-446655440000",
		Name:         "BE 17",
		StartDate:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		GraduateDate: time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		CreatedAt:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		UpdatedAt:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		UserID:       "65d8040e-03ae-11ee-88e5-e8fb1c216033",
		PIC:          "user1",
	}

	result := class.ClassCore{
		ClassID:      "550e8400-e29b-41d4-a716-446655440000",
		Name:         "BE 17",
		StartDate:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		GraduateDate: time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		CreatedAt:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		UpdatedAt:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		UserID:       "65d8040e-03ae-11ee-88e5-e8fb1c216033",
		PIC:          "user1",
	}

	service := New(data)

	t.Run("request cannot be empty", func(t *testing.T) {
		request := class.ClassCore{
			Name:         "BE 17",
			StartDate:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
			GraduateDate: time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
			UserID:       "",
		}
		_, _, err := service.RegisterClass(request)
		expectedErr := errors.New("request cannot be empty")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("success create a class", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(result, nil).Once()
		res, _, err := service.RegisterClass(arguments)
		assert.Nil(t, err)
		assert.Equal(t, result.UserID, res.UserID)
		assert.NotEmpty(t, result.ClassID)
		assert.NotEmpty(t, result.Name)
		assert.NotEmpty(t, result.StartDate)
		assert.NotEmpty(t, result.GraduateDate)
		data.AssertExpectations(t)
	})

	t.Run("error insert data, duplicated", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(class.ClassCore{}, errors.New("error insert data, duplicated")).Once()
		res, _, err := service.RegisterClass(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.ClassID)
		assert.ErrorContains(t, err, "duplicated")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(class.ClassCore{}, errors.New("internal server error")).Once()
		res, _, err := service.RegisterClass(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.ClassID)
		assert.ErrorContains(t, err, "internal server error")
		data.AssertExpectations(t)
	})
}

func TestListClasses(t *testing.T) {
	data := mocks.NewClassData(t)
	service := New(data)
	limit := 5
	offset := 0

	expectedResult := []class.ClassCore{
		{ClassID: "550e8400-e29b-41d4-a716-446655440000", Name: "FE 14"},
	}
	expectedCount := uint(1)

	t.Run("success", func(t *testing.T) {

		data.On("ListClasses", limit, offset).Return(expectedResult, expectedCount, nil)

		result, _, err := service.ListClasses(limit, offset)

		assert.Nil(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, expectedResult[0].ClassID, result[0].ClassID)
		assert.Equal(t, expectedResult[0].Name, result[0].Name)
		data.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		data.On("ListClasses", limit, offset).Return([]class.ClassCore{}, uint(0), errors.New("not found, error while retreiving list classes")).Once()

		result, _, err := service.ListClasses(limit, offset)

		assert.NotNil(t, uint(0), err)
		assert.Empty(t, result)
		assert.EqualError(t, err, "not found, error while retreiving list classes")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("ListClasses", limit, offset).Return([]class.ClassCore{}, uint(0), errors.New("internal server error")).Once()

		result, _, err := service.ListClasses(limit, offset)

		assert.NotNil(t, uint(0), err)
		assert.Empty(t, result)
		assert.EqualError(t, err, "internal server error")
		data.AssertExpectations(t)
	})

}

func TestDeleteClass(t *testing.T) {
	data := mocks.NewClassData(t)
	service := New(data)

	t.Run("success deactivate an account", func(t *testing.T) {
		data.On("DeleteClass", "550e8400-e29b-41d4-a716-446655440000").Return(nil).Once()
		err := service.DeleteClass("550e8400-e29b-41d4-a716-446655440000")
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("class record not found", func(t *testing.T) {
		data.On("DeleteClass", "550e8400-e29b-41d4-a716-446655440000").Return(errors.New("class record not found")).Once()
		err := service.DeleteClass("550e8400-e29b-41d4-a716-446655440000")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "class record not found")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("DeleteClass", "550e8400-e29b-41d4-a716-446655440000").Return(errors.New("internal server error")).Once()
		err := service.DeleteClass("550e8400-e29b-41d4-a716-446655440000")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		data.AssertExpectations(t)
	})
}

func TestGetClass(t *testing.T) {
	data := mocks.NewClassData(t)
	result := class.ClassCore{
		ClassID:      "550e8400-e29b-41d4-a716-446655440000",
		Name:         "BE 17",
		StartDate:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		GraduateDate: time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		CreatedAt:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		UpdatedAt:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		UserID:       "65d8040e-03ae-11ee-88e5-e8fb1c216033",
		PIC:          "user1",
	}
	service := New(data)

	t.Run("success get profile", func(t *testing.T) {
		data.On("GetClass", "550e8400-e29b-41d4-a716-446655440000").Return(result, nil).Once()
		res, err := service.GetClass("550e8400-e29b-41d4-a716-446655440000")
		assert.Nil(t, err)
		assert.Equal(t, result.ClassID, res.ClassID)
		assert.Equal(t, result.Name, res.Name)
		assert.Equal(t, result.PIC, res.PIC)
		data.AssertExpectations(t)
	})

	t.Run("not found, error while retrieving class", func(t *testing.T) {
		data.On("GetClass", "550e8400-e29b-41d4-a716-446655440000").Return(class.ClassCore{}, errors.New("not found, error while retrieving class")).Once()
		res, err := service.GetClass("550e8400-e29b-41d4-a716-446655440000")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, class.ClassCore{}, res)
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("GetClass", "550e8400-e29b-41d4-a716-446655440000").Return(class.ClassCore{}, errors.New("internal server error")).Once()
		res, err := service.GetClass("550e8400-e29b-41d4-a716-446655440000")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		assert.Equal(t, class.ClassCore{}, res)
		data.AssertExpectations(t)
	})
}

func TestUpdateProfile(t *testing.T) {
	data := mocks.NewClassData(t)
	arguments := class.ClassCore{
		Name:         "BE 17",
		PIC:          "550e8400-ffff-41d4-a716-446655440000",
		StartDate:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		GraduateDate: time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
	}
	result := class.ClassCore{
		ClassID:      "550e8400-e29b-41d4-a716-446655440000",
		Name:         "BE 17",
		PIC:          "550e8400-ffff-41d4-a716-446655440000",
		StartDate:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		GraduateDate: time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
	}
	service := New(data)

	t.Run("request cannot be empty", func(t *testing.T) {
		request := class.ClassCore{
			ClassID: "",
			Name:    "",
			PIC:     "",
		}
		err := service.UpdateClass("550e8400-e29b-41d4-a716-446655440000", request)
		expectedErr := errors.New("request cannot be empty")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("success update account", func(t *testing.T) {
		data.On("UpdateProfile", mock.Anything).Return(result, nil).Once()
		err := service.UpdateClass("550e8400-e29b-41d4-a716-446655440000", arguments)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("class record not found", func(t *testing.T) {
		data.On("UpdateProfile", "550e8400-e29b-41d4-a716-446655440000").Return(errors.New("class record not found")).Once()
		err := service.UpdateClass("550e8400-e29b-41d4-a716-446655440000", arguments)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "class record not found")
		data.AssertExpectations(t)
	})

	t.Run("failed to update class, duplicate data entry", func(t *testing.T) {
		data.On("UpdateProfile", "550e8400-e29b-41d4-a716-446655440000", arguments).Return(errors.New("failed to update class, duplicate data entry")).Once()
		err := service.UpdateClass("550e8400-e29b-41d4-a716-446655440000", arguments)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "duplicated")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("UpdateProfile", "550e8400-e29b-41d4-a716-446655440000").Return(errors.New("internal server error")).Once()
		err := service.UpdateClass("550e8400-e29b-41d4-a716-446655440000", arguments)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		data.AssertExpectations(t)
	})
}
