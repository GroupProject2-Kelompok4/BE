package service

import (
	"errors"
	"testing"
	"time"

	"github.com/GroupProject2-Kelompok4/BE/features/class"
	"github.com/GroupProject2-Kelompok4/BE/features/mentee"
	"github.com/GroupProject2-Kelompok4/BE/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterClass(t *testing.T) {
	data := mocks.NewMenteeData(t)
	arguments := mentee.MenteeCore{
		Fullname:        "zalai",
		Nickname:        "zalai",
		Phone:           "2313424",
		Email:           "zalai@example.com",
		CurrentAddress:  "123 Main Street",
		HomeAddress:     "456 Elm Street",
		Telegram:        "zalai123",
		Gender:          "M",
		EducationType:   "Informatics",
		Major:           "Computer Science",
		GraduateDate:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		Institution:     "University of Example",
		EmergencyName:   "janesmith",
		EmergencyPhone:  "0987654321",
		EmergencyStatus: "mother",
		Status:          "active",
		ClassID:         "8335bbaa-047a-11ee-8cfa-e8fb1c216033",
	}

	result := mentee.MenteeCore{
		MenteeID:        "7fcc04eb-047b-11ee-a7ae-e8fb1c216033",
		Fullname:        "zalai",
		Nickname:        "zalai",
		Phone:           "2313424",
		Email:           "zalai@example.com",
		CurrentAddress:  "123 Main Street",
		HomeAddress:     "456 Elm Street",
		Telegram:        "zalai123",
		Gender:          "M",
		EducationType:   "Informatics",
		Major:           "Computer Science",
		GraduateDate:    time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		Institution:     "University of Example",
		EmergencyName:   "janesmith",
		EmergencyPhone:  "0987654321",
		EmergencyStatus: "mother",
		Status:          "active",
		ClassID:         "8335bbaa-047a-11ee-8cfa-e8fb1c216033",
	}

	service := New(data)

	t.Run("request cannot be empty", func(t *testing.T) {
		request := mentee.MenteeCore{
			Fullname: "",
			Nickname: "",
			Phone:    "",
			ClassID:  "",
		}
		_, err := service.RegisterMentee(request)
		expectedErr := errors.New("request cannot be empty")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("success create a class", func(t *testing.T) {
		data.On("RegisterMentee", mock.Anything).Return(result, nil).Once()
		res, err := service.RegisterMentee(arguments)
		assert.Nil(t, err)
		assert.Equal(t, result.MenteeID, res.MenteeID)
		assert.NotEmpty(t, result.ClassID)
		assert.NotEmpty(t, result.Fullname)
		assert.NotEmpty(t, result.Nickname)
		assert.NotEmpty(t, result.Phone)
		data.AssertExpectations(t)
	})

	t.Run("error insert data, duplicated", func(t *testing.T) {
		data.On("RegisterMentee", mock.Anything).Return(class.ClassCore{}, errors.New("error insert data, duplicated")).Once()
		res, err := service.RegisterMentee(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.MenteeID)
		assert.ErrorContains(t, err, "duplicated")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("RegisterMentee", mock.Anything).Return(class.ClassCore{}, errors.New("internal server error")).Once()
		res, err := service.RegisterMentee(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.MenteeID)
		assert.ErrorContains(t, err, "internal server error")
		data.AssertExpectations(t)
	})
}
