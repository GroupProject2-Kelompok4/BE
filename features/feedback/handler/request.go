package handler

import "github.com/GroupProject2-Kelompok4/BE/features/feedback"

type RegisterFeedbackMenteeRequest struct {
	Notes    string `json:"notes" form:"notes"`
	Proof    string `json:"proof" form:"proof"`
	MenteeID string `json:"mentee_id" form:"mentee_id"`
}

type UpdateFeedbackMenteeRequest struct {
	Notes    *string `json:"notes" form:"notes"`
	Proof    *string `json:"proof" form:"proof"`
	MenteeID *string `json:"mentee_id" form:"mentee_id"`
}

func RequestToCore(data interface{}) feedback.FeedbackCore {
	res := feedback.FeedbackCore{}
	switch v := data.(type) {
	case RegisterFeedbackMenteeRequest:
		res.Notes = v.Notes
		res.Proof = v.Proof
		res.MenteeID = v.MenteeID
	case *UpdateFeedbackMenteeRequest:
		if v.Notes != nil {
			res.Notes = *v.Notes
		}
		if v.Proof != nil {
			res.Notes = *v.Proof
		}
		if v.MenteeID != nil {
			res.Notes = *v.MenteeID
		}
	default:
		return feedback.FeedbackCore{}
	}
	return res
}
