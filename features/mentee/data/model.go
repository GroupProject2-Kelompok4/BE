package data

import (
	"time"

	feedback "github.com/GroupProject2-Kelompok4/BE/features/feedback/data"
)

type Mentee struct {
	MenteeID        string    `gorm:"primaryKey;type:varchar(50)"`
	Fullname        string    `gorm:"type:varchar(100);not null"`
	Nickname        string    `gorm:"type:varchar(100);not null"`
	Email           string    `gorm:"type:varchar(100);not null;unique"`
	CurrentAddress  string    `gorm:"type:varchar(255)"`
	HomeAddress     string    `gorm:"type:varchar(255)"`
	Telegram        string    `gorm:"type:varchar(50);not null;unique"`
	Gender          string    `gorm:"type:enum('M', 'F');default:'M'"`
	EducationType   string    `gorm:"type:enum('Informatics', 'Non-Informatics');default:'Informatics'"`
	Major           string    `gorm:"type:varchar(255)"`
	GraduateDate    time.Time `gorm:"type:date"`
	Institution     string    `gorm:"type:varchar(255)"`
	EmergencyName   string    `gorm:"type:varchar(100)"`
	EmergencyPhone  string    `gorm:"type:varchar(15)"`
	EmergencyStatus string    `gorm:"type:varchar(100)"`
	CreatedAt       time.Time `gorm:"type:datetime"`
	UpdatedAt       time.Time `gorm:"type:datetime"`
	Status          string    `gorm:"type:varchar(100)"`
	ClassID         string
	Feedbacks       []feedback.Feedback `gorm:"foreignKey:MenteeID"`
}
