package data

import (
	"errors"

	"github.com/GroupProject2-Kelompok4/BE/features/mentee"
	"github.com/GroupProject2-Kelompok4/BE/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var log = utils.Log()

type menteeQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.MenteeData {
	return &menteeQuery{
		db: db,
	}
}

// RegisterMentee implements mentee.MenteeData
func (mq *menteeQuery) RegisterMentee(request mentee.MenteeCore) (mentee.MenteeCore, error) {
	menteeId, err := uuid.NewUUID()
	if err != nil {
		log.Warn("error while create uuid for class")
		return mentee.MenteeCore{}, nil
	}

	request.MenteeID = menteeId.String()
	req := menteeEntities(request)
	query := mq.db.Table("mentees").Create(&req)
	if query.Error != nil {
		log.Error("error insert data, duplicated")
		return mentee.MenteeCore{}, errors.New("error insert data, duplicated")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		log.Warn("no user has been created")
		return mentee.MenteeCore{}, errors.New("row affected : 0")
	}

	return menteeModels(req), nil
}

// SearchMentee implements mentee.MenteeData
func (mq *menteeQuery) SearchMentee(keyword string, limit int, offset int) ([]mentee.MenteeCore, uint, error) {
	mentees := []Mentee{}
	var count int64
	query := mq.db.Table("mentees").
		Select("mentees.*, classes.name").
		Joins("JOIN classes ON mentees.class_id = classes.class_id").
		Where("mentees.status LIKE ? OR mentees.education_type LIKE ? OR classes.name LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").
		Where("mentees.is_deleted = 0").
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Preload("Class").
		Find(&mentees).
		Count(&count)

	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("list mentees not found")
		return nil, 0, errors.New("mentees not found")
	}

	result := make([]mentee.MenteeCore, len(mentees))
	for i, mentee := range mentees {
		result[i] = menteeModels(mentee)
	}

	return result, uint(count), nil
}

// ProfileMenteeAndFeedback implements mentee.MenteeData
func (mq *menteeQuery) ProfileMenteeAndFeedback(menteeId string) (mentee.MenteeCore, error) {
	menteeLog := Mentee{}
	query := mq.db.Table("mentees").
		Select("mentees.*, users.fullname").
		Joins("JOIN feedbacks ON mentees.mentee_id = feedbacks.mentee_id").
		Joins("JOIN users ON mentees.user_id = users.user_id").
		Preload("User").
		Preload("Feedbacks").
		Where("mentees.mentee_id = ? AND mentees.is_deleted = 0", menteeId).
		First(&menteeLog)
	if query.Error != nil {
		log.Error("failed to fetch mentee profile and feedbacks")
		return mentee.MenteeCore{}, query.Error
	}

	return modeltoCore(menteeLog), nil
}

// UpdateMentee implements mentee.MenteeData
func (mq *menteeQuery) UpdateMentee(menteeId string, request mentee.MenteeCore) error {
	query := mq.db.Table("mentees").Where("mentee_id = ? AND is_deleted = 0", menteeId).Updates(menteeEntities(request))
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("mentee profile record not found")
		return errors.New("mentee profile record not found")
	}

	if query.RowsAffected == 0 {
		log.Warn("no mentee has been created")
		return errors.New("row affected : 0")
	}

	if query.Error != nil {
		log.Error("error while updating mentee")
		return errors.New("duplicate data entry")
	}

	return nil
}

// DeleteMentee implements mentee.MenteeData
func (mq *menteeQuery) DeactiveMentee(menteeId string) error {
	query := mq.db.Table("mentees").Where("mentee_id = ? AND is_deleted = 0", menteeId).Updates(map[string]interface{}{
		"is_deleted": true,
		"status":     "deleted",
	})

	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("mentee profile record not found")
		return errors.New("mentee profile record not found")
	}

	if query.RowsAffected == 0 {
		log.Warn("no mentee has been created")
		return errors.New("row affected : 0")
	}

	if query.Error != nil {
		log.Error("error while deactivate mentee")
		return errors.New("duplicate data entry")
	}

	return nil
}

// ProfileMentee implements mentee.MenteeData
func (mq *menteeQuery) ProfileMentee(menteeId string) (mentee.MenteeCore, error) {
	menteeLog := Mentee{}
	query := mq.db.Table("mentees").
		Select("mentees.*").
		Joins("JOIN feedbacks ON mentees.mentee_id = feedbacks.mentee_id").
		Joins("JOIN users AS feedback_users ON feedbacks.user_id = feedback_users.user_id").
		Joins("JOIN mentees AS feedback_mentees ON feedbacks.mentee_id = feedback_mentees.mentee_id").
		Preload("Class").
		Preload("Feedbacks").
		Preload("Feedbacks.User").
		Preload("Feedbacks.Mentee").
		Where("mentees.mentee_id = ? AND mentees.is_deleted = 0", menteeId).
		First(&menteeLog)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("mentee profile record not found")
		return mentee.MenteeCore{}, errors.New("mentee profile record not found")
	}

	return modeltoCore(menteeLog), nil
}
