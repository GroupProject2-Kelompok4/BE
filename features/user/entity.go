package user

import (
	"time"

	class "github.com/GroupProject2-Kelompok4/BE/features/class"
	feedback "github.com/GroupProject2-Kelompok4/BE/features/feedback"
	"github.com/labstack/echo/v4"
)

type UserCore struct {
	No          uint
	UserID      string
	Fullname    string
	Email       string
	Password    string
	Team        string
	Role        string
	Status      string
	UserPicture string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool
	Classes     []class.ClassCore
	Feedbacks   []feedback.FeedbackCore
}

type UserHandler interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	SearchUser() echo.HandlerFunc
}

type UserService interface {
	Login(request UserCore) (UserCore, string, error)
	Register(request UserCore) (UserCore, error)
	SearchUser(keyword string, limit, offset int) ([]UserCore, uint, error)
}

type UserData interface {
	Login(request UserCore) (UserCore, string, error)
	Register(request UserCore) (UserCore, error)
	SearchUser(keyword string, limit, offset int) ([]UserCore, uint, error)
}
