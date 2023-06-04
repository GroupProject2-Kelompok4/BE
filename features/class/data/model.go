package data

import (
	"time"

	mentee "github.com/GroupProject2-Kelompok4/BE/features/mentee/data"
)

type Class struct {
	ClassID      string    `gorm:"primaryKey;type:varchar(50)"`
	Name         string    `gorm:"type:varchar(5);not null;unique"`
	StartDate    time.Time `gorm:"type:date"`
	GraduateDate time.Time `gorm:"type:date"`
	CreatedAt    time.Time `gorm:"type:datetime"`
	UpdatedAt    time.Time `gorm:"type:datetime"`
	UserID       string
	Mentees      []mentee.Mentee `gorm:"foreignKey:ClassID"`
}
