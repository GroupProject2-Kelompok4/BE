package service

import (
	"errors"
	"strings"

	"github.com/GroupProject2-Kelompok4/BE/features/feedback"
	"github.com/GroupProject2-Kelompok4/BE/utils"
)

var log = utils.Log()

type feedbackService struct {
	query feedback.FeedbackData
}

func New(fd feedback.FeedbackData) feedback.FeedbackService {
	return &feedbackService{
		query: fd,
	}
}

// RegisterFeedback implements feedback.FeedbackService
func (fs *feedbackService) RegisterFeedbackMentee(request feedback.FeedbackCore, userId string) (feedback.FeedbackCore, error) {
	if request.Notes == "" || request.MenteeID == "" {
		log.Error("request cannot be empty")
		return feedback.FeedbackCore{}, errors.New("request cannot be empty")
	}

	result, err := fs.query.RegisterFeedbackMentee(request, userId)
	if err != nil {
		log.Error("internal server error")
		return feedback.FeedbackCore{}, errors.New("internal server error")
	}

	return result, nil
}

// UpdateFeedbackMentee implements feedback.FeedbackService
func (fs *feedbackService) UpdateFeedbackMentee(request feedback.FeedbackCore, feedbackId, userId string) error {
	err := fs.query.UpdateFeedbackMentee(request, feedbackId, userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("feedback record not found")
			return errors.New("feedbackrecord not found")
		} else if strings.Contains(err.Error(), "duplicate data entry") {
			log.Error("failed to update feedback, duplicate data entry")
			return errors.New("failed to update feedback, duplicate data entry")
		} else {
			log.Error("internal server error")
			return errors.New("internal server error")
		}
	}

	return nil
}
