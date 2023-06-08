package data

import (
	"errors"

	"github.com/GroupProject2-Kelompok4/BE/features/feedback"
	"github.com/GroupProject2-Kelompok4/BE/utils"
	"gorm.io/gorm"
)

var log = utils.Log()

type feedbackQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedback.FeedbackData {
	return &feedbackQuery{
		db: db,
	}
}

// RegisterFeedbackMentee implements feedback.FeedbackData
func (fq *feedbackQuery) RegisterFeedbackMentee(request feedback.FeedbackCore, userId string) (feedback.FeedbackCore, error) {
	request.UserID = userId
	req := feedbackEntities(request)
	query := fq.db.Table("feedbacks").Create(&req)
	if query.Error != nil {
		log.Error("error inserting data")
		return feedback.FeedbackCore{}, query.Error
	}

	if query.RowsAffected == 0 {
		log.Warn("no feedback has been registered")
		return feedback.FeedbackCore{}, errors.New("no feedback has been registered")
	}

	var resp Feedback
	query = fq.db.Table("feedbacks").
		Select("feedbacks.*, users.fullname, mentees.status").
		Joins("JOIN users ON users.user_id = feedbacks.user_id").
		Joins("JOIN mentees ON mentees.mentee_id = feedbacks.mentee_id").
		Preload("User").
		Preload("Mentee").
		Last(&resp)
	if query.Error != nil {
		log.Error("error while retrieving feedback")
		return feedback.FeedbackCore{}, errors.New("error while retrieving feedback")
	}

	return feedbackModels(resp), nil
}

// UpdateFeedbackMentee implements feedback.FeedbackData
func (fq *feedbackQuery) UpdateFeedbackMentee(request feedback.FeedbackCore, feedbackId, userId string) error {
	req := feedbackEntities(request)
	query := fq.db.Table("feedbacks").Where("feedback_id = ? AND user_id = ? AND is_deleted = 0", feedbackId, userId).Updates(&req)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("feedback record not found")
		return errors.New("feedback record not found")
	}

	if query.RowsAffected == 0 {
		log.Warn("no feedback has been created")
		return errors.New("row affected : 0")
	}

	if query.Error != nil {
		log.Error("error while updating feedback")
		return errors.New("duplicate data entry")
	}

	return nil
}

// DeleteFeedbackMentee implements feedback.FeedbackData
func (fq *feedbackQuery) DeleteFeedbackMentee(feedbackId string, userId string) error {
	query := fq.db.Table("feedbacks").Where("feedback_id = ? AND user_id = ? AND is_deleted = 0", feedbackId, userId).Updates(map[string]interface{}{
		"is_deleted": true,
	})

	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("feedbackrecord not found")
		return errors.New("feedback record not found")
	}

	if query.RowsAffected == 0 {
		log.Warn("no feedback has been created")
		return errors.New("row affected : 0")
	}

	if query.Error != nil {
		log.Error("error while deleted feedback")
		return errors.New("internal server error")
	}

	return nil
}
