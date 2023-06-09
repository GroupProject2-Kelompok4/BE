package service

import (
	"errors"
	"strings"

	"github.com/GroupProject2-Kelompok4/BE/features/class"
	"github.com/GroupProject2-Kelompok4/BE/utils"
)

var log = utils.Log()

type classService struct {
	query class.ClassData
}

func New(cd class.ClassData) class.ClassService {
	return &classService{
		query: cd,
	}
}

// RegisterClass implements class.ClassService
func (cs *classService) RegisterClass(request class.ClassCore) (class.ClassCore, string, error) {
	if request.Name == "" || request.UserID == "" || request.StartDate.IsZero() || request.GraduateDate.IsZero() {
		log.Error("request cannot be empty")
		return class.ClassCore{}, "", errors.New("request cannot be empty")
	}

	result, pic, err := cs.query.RegisterClass(request)
	if err != nil {
		message := ""
		if strings.Contains(err.Error(), "error insert data, duplicated") {
			log.Error("error insert data, duplicated")
			message = "error insert data, duplicated"
		} else {
			log.Error("internal server error")
			message = "internal server error"
		}
		return class.ClassCore{}, "", errors.New(message)
	}

	return result, pic, nil
}

// ListClasses implements class.ClassService
func (cs *classService) ListClasses(limit int, offset int) ([]class.ClassCore, uint, error) {
	result, count, err := cs.query.ListClasses(limit, offset)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("not found, error while retreiving list classes")
			return []class.ClassCore{}, 0, errors.New("not found, error while retreiving list classes")
		} else {
			log.Error("internal server error")
			return []class.ClassCore{}, 0, errors.New("internal server error")
		}
	}
	return result, count, nil
}

// DeleteClass implements class.ClassService
func (cs *classService) DeleteClass(classId string) error {
	err := cs.query.DeleteClass(classId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("class record not found")
			return errors.New("class record not found")
		} else {
			log.Error("internal server error")
			return errors.New("internal server error")
		}
	}

	return nil
}

// GetClass implements class.ClassService
func (cs *classService) GetClass(classId string) (class.ClassCore, error) {
	result, err := cs.query.GetClass(classId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("not found, error while retrieving class")
			return class.ClassCore{}, errors.New("not found, error while retrieving class")
		} else {
			log.Error("internal server error")
			return class.ClassCore{}, errors.New("internal server error")
		}
	}
	return result, nil
}

// UpdateClass implements class.ClassService
func (cs *classService) UpdateClass(classId string, request class.ClassCore) error {
	if request.Name == "" && request.UserID == "" && request.StartDate.IsZero() && request.GraduateDate.IsZero() {
		log.Error("request cannot be empty")
		return errors.New("request cannot be empty")
	}

	err := cs.query.UpdateClass(classId, request)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("class record not found")
			return errors.New("class record not found")
		} else if strings.Contains(err.Error(), "duplicate data entry") {
			log.Error("failed to update class, duplicate data entry")
			return errors.New("failed to update class, duplicate data entry")
		} else {
			log.Error("internal server error")
			return errors.New("internal server error")
		}
	}

	return nil
}
