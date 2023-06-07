package handler

import "github.com/GroupProject2-Kelompok4/BE/features/feedback"

type RegisterFeedbackMenteeRequest struct {
	Notes    string `json:"notes" form:"notes"`
	Proof    string `json:"proof" form:"proof"`
	MenteeID string `json:"mentee_id" form:"mentee_id"`
}

func RequestToCore(data interface{}) feedback.FeedbackCore {
	res := feedback.FeedbackCore{}
	switch v := data.(type) {
	case RegisterFeedbackMenteeRequest:
		res.Notes = v.Notes
		res.Proof = v.Proof
		res.MenteeID = v.MenteeID
	default:
		return feedback.FeedbackCore{}
	}
	return res
}
