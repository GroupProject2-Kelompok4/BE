package data

import (
	"time"

	class "github.com/GroupProject2-Kelompok4/BE/features/class/data"
	feedback "github.com/GroupProject2-Kelompok4/BE/features/feedback/data"
	mentee "github.com/GroupProject2-Kelompok4/BE/features/mentee/data"
	"github.com/GroupProject2-Kelompok4/BE/features/user"
)

type User struct {
	UserID      string `gorm:"primaryKey;type:varchar(50)"`
	Fullname    string `gorm:"type:varchar(100);not null;unique"`
	Email       string `gorm:"type:varchar(100);not null;unique"`
	Password    string
	Team        string `gorm:"type:enum('manager', 'mentor', 'team people skill', 'team placement'); default:'mentor'"`
	Role        string `gorm:"type:enum('admin', 'user'); default:'user'"`
	Status      string `gorm:"type:enum('active', 'non-active', 'deleted'); default:'active'"`
	UserPicture string
	CreatedAt   time.Time           `gorm:"type:datetime"`
	UpdatedAt   time.Time           `gorm:"type:datetime"`
	IsDeleted   bool                `gorm:"type:boolean"`
	Mentees     []mentee.Mentee     `gorm:"foreignKey:UserID"`
	Classes     []class.Class       `gorm:"foreignKey:UserID"`
	Feedbacks   []feedback.Feedback `gorm:"foreignKey:UserID"`
}

// User-model to user-core
func userModels(u User) user.UserCore {
	return user.UserCore{
		UserID:      u.UserID,
		Fullname:    u.Fullname,
		Email:       u.Email,
		Password:    u.Password,
		Team:        u.Team,
		Role:        u.Role,
		Status:      u.Status,
		UserPicture: u.UserPicture,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		IsDeleted:   u.IsDeleted,
	}
}

// User-core to user-model
func userEntities(u user.UserCore) User {
	return User{
		UserID:      u.UserID,
		Fullname:    u.Fullname,
		Email:       u.Email,
		Password:    u.Password,
		Team:        u.Team,
		Role:        u.Role,
		Status:      u.Status,
		UserPicture: u.UserPicture,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		IsDeleted:   u.IsDeleted,
	}
}
