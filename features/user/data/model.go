package data

import (
	"time"

	class "github.com/GroupProject2-Kelompok4/BE/features/class/data"
	feedback "github.com/GroupProject2-Kelompok4/BE/features/feedback/data"
	"github.com/GroupProject2-Kelompok4/BE/features/user"
	"github.com/google/uuid"
)

type User struct {
	UserID      string `gorm:"primaryKey;type:varchar(50)"`
	Fullname    string `gorm:"type:varchar(100);not null;unique"`
	Email       string `gorm:"type:varchar(100);not null;unique"`
	Password    string
	Role        string `gorm:"type:enum('admin', 'user');default:'user'"`
	Status      string `gorm:"type:enum('manager', 'mentor', 'team people skill', 'team placement'); default:'mentor'"`
	UserPicture string
	CreatedAt   time.Time           `gorm:"type:datetime"`
	UpdatedAt   time.Time           `gorm:"type:datetime"`
	IsDeleted   bool                `gorm:"type:boolean"`
	Classes     []class.Class       `gorm:"foreignKey:UserID"`
	Feedbacks   []feedback.Feedback `gorm:"foreignKey:UserID"`
}

// User-model to user-core
func userModels(u User) user.UserCore {
	User_ID, _ := uuid.NewUUID()
	return user.UserCore{
		UserID:      User_ID.String(),
		Fullname:    u.Fullname,
		Email:       u.Email,
		Password:    u.Password,
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
	User_ID, _ := uuid.NewUUID()
	return User{
		UserID:      User_ID.String(),
		Fullname:    u.Fullname,
		Email:       u.Email,
		Password:    u.Password,
		Role:        u.Role,
		Status:      u.Status,
		UserPicture: u.UserPicture,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		IsDeleted:   u.IsDeleted,
	}
}
