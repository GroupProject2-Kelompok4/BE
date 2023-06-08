package feedback

import (
	"time"

	"github.com/labstack/echo/v4"
)

type FeedbackCore struct {
	FeedbackID uint
	Notes      string
	Proof      string
	Approved   bool
	Status     string
	Users      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     string
	MenteeID   string
}

type FeedbackHandler interface {
	RegisterFeedbackMentee() echo.HandlerFunc
	UpdateFeedbackMentee() echo.HandlerFunc
}

type FeedbackService interface {
	RegisterFeedbackMentee(request FeedbackCore, userId string) (FeedbackCore, error)
	UpdateFeedbackMentee(request FeedbackCore, feedbackId, userId string) error
}

type FeedbackData interface {
	RegisterFeedbackMentee(request FeedbackCore, userId string) (FeedbackCore, error)
	UpdateFeedbackMentee(request FeedbackCore, feedbackId, userId string) error
}
