package data

import (
	"time"

	"github.com/GroupProject2-Kelompok4/BE/features/class"
	mentee "github.com/GroupProject2-Kelompok4/BE/features/mentee/data"
)

type Class struct {
	ClassID      string          `gorm:"primaryKey;type:varchar(50)"`
	Name         string          `gorm:"type:varchar(100);not null;unique"`
	StartDate    time.Time       `gorm:"type:date"`
	GraduateDate time.Time       `gorm:"type:date"`
	CreatedAt    time.Time       `gorm:"type:datetime"`
	UpdatedAt    time.Time       `gorm:"type:datetime"`
	IsDeleted    bool            `gorm:"type:boolean"`
	UserID       string          `gorm:"type:varchar(50)"`
	User         User            `gorm:"references:UserID"`
	Mentees      []mentee.Mentee `gorm:"foreignKey:ClassID"`
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
	CreatedAt   time.Time `gorm:"type:datetime"`
	UpdatedAt   time.Time `gorm:"type:datetime"`
	IsDeleted   bool      `gorm:"type:boolean"`
	Classes     []Class   `gorm:"foreignKey:UserID"`
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
		IsDeleted:    c.IsDeleted,
		UserID:       c.UserID,
		PIC:          c.User.Fullname,
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
		IsDeleted:    c.IsDeleted,
		UserID:       c.UserID,
	}
}
