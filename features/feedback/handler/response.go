package handler

import "github.com/GroupProject2-Kelompok4/BE/features/feedback"

type registerFeedbackMenteeResponse struct {
	FeedbackID uint
	Notes      string
	Users      string
	Status     string
	Proof      string
}

func registerFeedbackMentee(m feedback.FeedbackCore) registerFeedbackMenteeResponse {
	return registerFeedbackMenteeResponse{
		FeedbackID: m.FeedbackID,
		Notes:      m.Notes,
		Users:      m.Users,
		Status:     m.Status,
		Proof:      m.Proof,
	}
}
