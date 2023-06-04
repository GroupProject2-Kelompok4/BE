package data

import "time"

type Feedback struct {
	FeedbackID uint      `gorm:"primaryKey"`
	Notes      string    `gorm:"type:text"`
	Proof      string    `gorm:"type:varchar(255)"`
	Approved   bool      `gorm:"type:boolean"`
	CreatedAt  time.Time `gorm:"type:datetime"`
	UpdatedAt  time.Time `gorm:"type:datetime"`
	UserID     string
	MenteeID   string
}
