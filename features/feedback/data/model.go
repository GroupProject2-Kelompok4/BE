package data

import (
	"time"

	"github.com/GroupProject2-Kelompok4/BE/features/feedback"
)

type Feedback struct {
	FeedbackID uint      `gorm:"primaryKey;autoIncrement"`
	Notes      string    `gorm:"type:text"`
	Proof      string    `gorm:"type:varchar(255)"`
	Approved   bool      `gorm:"type:boolean"`
	CreatedAt  time.Time `gorm:"type:datetime"`
	UpdatedAt  time.Time `gorm:"type:datetime"`
	IsDeleted  bool      `gorm:"type:boolean"`
	UserID     string    `gorm:"type:varchar(50)"`
	MenteeID   string    `gorm:"type:varchar(50)"`
	User       User      `gorm:"references:UserID"`
	Mentee     Mentee    `gorm:"references:MenteeID"`
}

type Mentee struct {
	MenteeID        string     `gorm:"primaryKey;type:varchar(50)"`
	Fullname        string     `gorm:"type:varchar(100);not null"`
	Nickname        string     `gorm:"type:varchar(100);not null"`
	Phone           string     `gorm:"type:varchar(15)"`
	Email           string     `gorm:"type:varchar(100);not null;unique"`
	CurrentAddress  string     `gorm:"type:varchar(255)"`
	HomeAddress     string     `gorm:"type:varchar(255)"`
	Telegram        string     `gorm:"type:varchar(50);not null;unique"`
	Gender          string     `gorm:"type:enum('M', 'F');default:'M'"`
	EducationType   string     `gorm:"type:enum('Informatics', 'Non-Informatics');default:'Informatics'"`
	Major           string     `gorm:"type:varchar(255)"`
	GraduateDate    time.Time  `gorm:"type:date"`
	Institution     string     `gorm:"type:varchar(255)"`
	EmergencyName   string     `gorm:"type:varchar(100)"`
	EmergencyPhone  string     `gorm:"type:varchar(15)"`
	EmergencyStatus string     `gorm:"type:varchar(100)"`
	CreatedAt       time.Time  `gorm:"type:datetime"`
	UpdatedAt       time.Time  `gorm:"type:datetime"`
	IsDeleted       bool       `gorm:"type:boolean"`
	Status          string     `gorm:"type:varchar(100)"`
	ClassID         string     `gorm:"type:varchar(50)"`
	Feedbacks       []Feedback `gorm:"foreignKey:MenteeID"`
}

type User struct {
	UserID      string `gorm:"primaryKey;type:varchar(50)"`
	Fullname    string `gorm:"type:varchar(100);not null;unique"`
	Email       string `gorm:"type:varchar(100);not null;unique"`
	Password    string
	Team        string `gorm:"type:enum('manager', 'mentor', 'team people skill', 'team placement'); default:'mentor'"`
	Role        string `gorm:"type:enum('admin', 'user'); default:'user'"`
	Status      string `gorm:"type:enum('active', 'non-active', 'deleted'); default:'active'"`
	UserPicture string
	CreatedAt   time.Time  `gorm:"type:datetime"`
	UpdatedAt   time.Time  `gorm:"type:datetime"`
	IsDeleted   bool       `gorm:"type:boolean"`
	Feedbacks   []Feedback `gorm:"foreignKey:UserID"`
}

func feedbackModels(f Feedback) feedback.FeedbackCore {
	return feedback.FeedbackCore{
		FeedbackID: f.FeedbackID,
		Notes:      f.Notes,
		Proof:      f.Proof,
		Status:     f.Mentee.Status,
		Users:      f.User.Fullname,
		Approved:   f.Approved,
		CreatedAt:  f.CreatedAt,
		UpdatedAt:  f.UpdatedAt,
		UserID:     f.UserID,
		MenteeID:   f.MenteeID,
	}
}

func feedbackEntities(f feedback.FeedbackCore) Feedback {
	return Feedback{
		FeedbackID: f.FeedbackID,
		Notes:      f.Notes,
		Proof:      f.Proof,
		Approved:   f.Approved,
		CreatedAt:  f.CreatedAt,
		UpdatedAt:  f.UpdatedAt,
		UserID:     f.UserID,
		MenteeID:   f.MenteeID,
	}
}
