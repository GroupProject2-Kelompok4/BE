package data

import (
	"time"

	feedback "github.com/GroupProject2-Kelompok4/BE/features/feedback/data"
	"github.com/GroupProject2-Kelompok4/BE/features/mentee"
)

type Mentee struct {
	MenteeID        string              `gorm:"primaryKey;type:varchar(50)"`
	Fullname        string              `gorm:"type:varchar(100);not null"`
	Nickname        string              `gorm:"type:varchar(100);not null"`
	Phone           string              `gorm:"type:varchar(15)"`
	Email           string              `gorm:"type:varchar(100);not null;unique"`
	CurrentAddress  string              `gorm:"type:varchar(255)"`
	HomeAddress     string              `gorm:"type:varchar(255)"`
	Telegram        string              `gorm:"type:varchar(50);not null;unique"`
	Gender          string              `gorm:"type:enum('M', 'F');default:'M'"`
	EducationType   string              `gorm:"type:enum('Informatics', 'Non-Informatics');default:'Informatics'"`
	Major           string              `gorm:"type:varchar(255)"`
	GraduateDate    time.Time           `gorm:"type:date"`
	Institution     string              `gorm:"type:varchar(255)"`
	EmergencyName   string              `gorm:"type:varchar(100)"`
	EmergencyPhone  string              `gorm:"type:varchar(15)"`
	EmergencyStatus string              `gorm:"type:varchar(100)"`
	CreatedAt       time.Time           `gorm:"type:datetime"`
	UpdatedAt       time.Time           `gorm:"type:datetime"`
	Status          string              `gorm:"type:varchar(100)"`
	ClassID         string              `gorm:"type:varchar(50)"`
	Class           Class               `gorm:"references:ClassID"`
	Feedbacks       []feedback.Feedback `gorm:"foreignKey:MenteeID"`
}

type Class struct {
	ClassID      string    `gorm:"primaryKey;type:varchar(50)"`
	Name         string    `gorm:"type:varchar(5);not null;unique"`
	StartDate    time.Time `gorm:"type:date"`
	GraduateDate time.Time `gorm:"type:date"`
	CreatedAt    time.Time `gorm:"type:datetime"`
	UpdatedAt    time.Time `gorm:"type:datetime"`
	IsDeleted    bool      `gorm:"type:boolean"`
	UserID       string    `gorm:"type:varchar(50)"`
	Mentees      []Mentee  `gorm:"foreignKey:ClassID"`
}

// Mentee-model to mentee-core
func menteeModels(m Mentee) mentee.MenteeCore {
	return mentee.MenteeCore{
		MenteeID:        m.MenteeID,
		Fullname:        m.Fullname,
		Nickname:        m.Nickname,
		Phone:           m.Phone,
		Email:           m.Email,
		CurrentAddress:  m.CurrentAddress,
		HomeAddress:     m.HomeAddress,
		Telegram:        m.Telegram,
		Gender:          m.Gender,
		EducationType:   m.EducationType,
		Major:           m.Major,
		GraduateDate:    m.GraduateDate,
		Institution:     m.Institution,
		EmergencyName:   m.EmergencyName,
		EmergencyPhone:  m.EmergencyPhone,
		EmergencyStatus: m.EmergencyStatus,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		Status:          m.Status,
		ClassID:         m.ClassID,
	}
}

// Mentee-core to mentee-model
func menteeEntities(m mentee.MenteeCore) Mentee {
	return Mentee{
		MenteeID:        m.MenteeID,
		Fullname:        m.Fullname,
		Nickname:        m.Nickname,
		Phone:           m.Phone,
		Email:           m.Email,
		CurrentAddress:  m.CurrentAddress,
		HomeAddress:     m.HomeAddress,
		Telegram:        m.Telegram,
		Gender:          m.Gender,
		EducationType:   m.EducationType,
		Major:           m.Major,
		GraduateDate:    m.GraduateDate,
		Institution:     m.Institution,
		EmergencyName:   m.EmergencyName,
		EmergencyPhone:  m.EmergencyPhone,
		EmergencyStatus: m.EmergencyStatus,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		Status:          m.Status,
		ClassID:         m.ClassID,
	}
}
