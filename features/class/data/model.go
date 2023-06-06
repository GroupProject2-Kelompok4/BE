package data

import (
	"time"

	"github.com/GroupProject2-Kelompok4/BE/features/class"
	mentee "github.com/GroupProject2-Kelompok4/BE/features/mentee/data"
)

type Class struct {
	ClassID      string          `gorm:"primaryKey;type:varchar(50)"`
	Name         string          `gorm:"type:varchar(5);not null;unique"`
	StartDate    time.Time       `gorm:"type:date"`
	GraduateDate time.Time       `gorm:"type:date"`
	CreatedAt    time.Time       `gorm:"type:datetime"`
	UpdatedAt    time.Time       `gorm:"type:datetime"`
	UserID       string          `gorm:"type:varchar(50)"`
	Mentees      []mentee.Mentee `gorm:"foreignKey:ClassID"`
}

// Class-model to class-core
func classModels(c Class) class.ClassCore {
	return class.ClassCore{
		ClassID:      c.ClassID,
		Name:         c.Name,
		StartDate:    c.StartDate,
		GraduateDate: c.GraduateDate,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
		UserID:       c.UserID,
	}
}

// Class-core to class-model
func classEntities(c class.ClassCore) Class {
	return Class{
		ClassID:      c.ClassID,
		Name:         c.Name,
		StartDate:    c.StartDate,
		GraduateDate: c.GraduateDate,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
		UserID:       c.UserID,
	}
}
