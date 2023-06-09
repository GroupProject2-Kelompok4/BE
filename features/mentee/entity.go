package mentee

import (
	"time"

	"github.com/labstack/echo/v4"
)

type MenteeCore struct {
	No              uint
	MenteeID        string
	Fullname        string
	Nickname        string
	Phone           string
	Email           string
	CurrentAddress  string
	HomeAddress     string
	Telegram        string
	Gender          string
	EducationType   string
	Major           string
	GraduateDate    time.Time
	Institution     string
	EmergencyName   string
	EmergencyPhone  string
	EmergencyStatus string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	IsDeleted       bool
	Status          string
	ClassID         string
	ClassName       string
	UserID          string
	Users           string
	Feedbacks       []FeedbackCore
}

type FeedbackCore struct {
	FeedbackId uint
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

type MenteeHandler interface {
	RegisterMentee() echo.HandlerFunc
	SearchMentee() echo.HandlerFunc
	ProfileMenteeAndFeedback() echo.HandlerFunc
	ProfileMentee() echo.HandlerFunc
	UpdateMentee() echo.HandlerFunc
	DeactiveMentee() echo.HandlerFunc
}

type MenteeService interface {
	RegisterMentee(request MenteeCore) (MenteeCore, error)
	SearchMentee(keyword string, limit int, offset int) ([]MenteeCore, uint, error)
	ProfileMenteeAndFeedback(menteeId string) (MenteeCore, error)
	ProfileMentee(menteeId string) (MenteeCore, error)
	UpdateMentee(menteeId string, request MenteeCore) error
	DeactiveMentee(menteeId string) error
}

type MenteeData interface {
	RegisterMentee(request MenteeCore) (MenteeCore, error)
	SearchMentee(keyword string, limit int, offset int) ([]MenteeCore, uint, error)
	ProfileMenteeAndFeedback(menteeId string) (MenteeCore, error)
	ProfileMentee(menteeId string) (MenteeCore, error)
	UpdateMentee(menteeId string, request MenteeCore) error
	DeactiveMentee(menteeId string) error
}
