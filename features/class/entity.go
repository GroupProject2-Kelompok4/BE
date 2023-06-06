package class

import (
	"time"

	mentee "github.com/GroupProject2-Kelompok4/BE/features/mentee"
	"github.com/labstack/echo/v4"
)

type ClassCore struct {
	No           uint
	ClassID      string
	Name         string
	StartDate    time.Time
	GraduateDate time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       string
	PIC          string
	Mentees      []mentee.MenteeCore
}

type ClassHandler interface {
	RegisterClass() echo.HandlerFunc
	ListClasses() echo.HandlerFunc
}

type ClassService interface {
	RegisterClass(request ClassCore) (ClassCore, string, error)
	ListClasses(limit, offset int) ([]ClassCore, uint, error)
}

type ClassData interface {
	RegisterClass(request ClassCore) (ClassCore, string, error)
	ListClasses(limit, offset int) ([]ClassCore, uint, error)
}
