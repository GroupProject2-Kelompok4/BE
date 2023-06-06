package handler

import "github.com/GroupProject2-Kelompok4/BE/features/mentee"

type registerMenteeResponse struct {
	MenteeID string `json:"mentee_id" form:"mentee_id"`
	Fullname string `json:"fullname" form:"fullname"`
	Nickname string `json:"nickname" form:"nickname"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Telegram string `json:"telegram" form:"telegram"`
	Gender   string `json:"gender" form:"gender"`
}

func registerMentee(m mentee.MenteeCore) registerMenteeResponse {
	return registerMenteeResponse{
		MenteeID: m.MenteeID,
		Fullname: m.Fullname,
		Nickname: m.Nickname,
		Email:    m.Email,
		Phone:    m.Phone,
		Telegram: m.Telegram,
		Gender:   m.Gender,
	}
}
