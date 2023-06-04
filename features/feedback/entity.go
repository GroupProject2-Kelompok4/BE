package feedback

import "time"

type FeedbackCore struct {
	FeedbackID uint
	Notes      string
	Proof      string
	Approved   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     string
	MenteeID   string
}
