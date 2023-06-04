package class

import (
	"time"

	mentee "github.com/GroupProject2-Kelompok4/BE/features/mentee"
)

type ClassCore struct {
	ClassID      string
	Name         string
	StartDate    time.Time
	GraduateDate time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       string
	Mentees      []mentee.MenteeCore
}
