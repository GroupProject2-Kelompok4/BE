package service

import (
	"errors"
	"strings"

	"github.com/GroupProject2-Kelompok4/BE/features/mentee"
	"github.com/GroupProject2-Kelompok4/BE/utils"
)

var log = utils.Log()

type menteeService struct {
	query mentee.MenteeData
}

func New(cd mentee.MenteeData) mentee.MenteeService {
	return &menteeService{
		query: cd,
	}
}

// RegisterMentee implements mentee.MenteeService
func (ms *menteeService) RegisterMentee(request mentee.MenteeCore) (mentee.MenteeCore, error) {
	if request.Fullname == "" || request.Nickname == "" || request.Phone == "" || request.ClassID == "" {
		log.Error("request cannot be empty")
		return mentee.MenteeCore{}, errors.New("request cannot be empty")
	}

	result, err := ms.query.RegisterMentee(request)
	if err != nil {
		message := ""
		if strings.Contains(err.Error(), "error insert data, duplicated") {
			log.Error("error insert data, duplicated")
			message = "error insert data, duplicated"
		} else {
			log.Error("internal server error")
			message = "internal server error"
		}
		return mentee.MenteeCore{}, errors.New(message)
	}

	return result, nil
}

// SearchMentee implements mentee.MenteeService
func (ms *menteeService) SearchMentee(keyword string, limit int, offset int) ([]mentee.MenteeCore, uint, error) {
	result, count, err := ms.query.SearchMentee(keyword, limit, offset)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("not found, error while retrieving list mentees")
			return []mentee.MenteeCore{}, 0, errors.New("not found, error while retrieving list mentees")
		} else {
			log.Error("internal server error")
			return []mentee.MenteeCore{}, 0, errors.New("internal server error")
		}
	}
	return result, count, nil
}

// ProfileMenteeAndFeedback implements mentee.MenteeService
func (ms *menteeService) ProfileMenteeAndFeedback(menteeId string) (mentee.MenteeCore, error) {
	result, err := ms.query.ProfileMenteeAndFeedback(menteeId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("not found, error while retrieving user profile")
			return mentee.MenteeCore{}, errors.New("not found, error while retrieving user profile")
		} else {
			log.Error("internal server error")
			return mentee.MenteeCore{}, errors.New("internal server error")
		}
	}
	return result, nil
}
