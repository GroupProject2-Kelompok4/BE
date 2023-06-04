package mentee

import (
	"time"

	feedback "github.com/GroupProject2-Kelompok4/BE/features/feedback"
)

type MenteeCore struct {
	MenteeID        string
	Fullname        string
	Nickname        string
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
