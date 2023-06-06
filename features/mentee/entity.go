package mentee

import (
	"time"

	feedback "github.com/GroupProject2-Kelompok4/BE/features/feedback"
	"github.com/labstack/echo/v4"
)

type MenteeCore struct {
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
	Status          string
	ClassID         string
	Feedbacks       []feedback.FeedbackCore
}

type MenteeHandler interface {
	RegisterMentee() echo.HandlerFunc
}

type MenteeService interface {
	RegisterMentee(request MenteeCore) (MenteeCore, error)
}

type MenteeData interface {
	RegisterMentee(request MenteeCore) (MenteeCore, error)
}
