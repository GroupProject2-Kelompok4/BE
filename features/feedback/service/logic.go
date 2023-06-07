package service

import (
	"errors"

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
func (fs *feedbackService) RegisterFeedbackMentee(request feedback.FeedbackCore) (feedback.FeedbackCore, error) {
	if request.Notes == "" || request.UserID == "" || request.MenteeID == "" {
		log.Error("request cannot be empty")
		return feedback.FeedbackCore{}, errors.New("request cannot be empty")
	}

	result, err := fs.query.RegisterFeedbackMentee(request)
	if err != nil {
		log.Error("internal server error")
		return feedback.FeedbackCore{}, errors.New("internal server error")
	}

	return result, nil
}
