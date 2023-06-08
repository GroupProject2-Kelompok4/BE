package handler

import (
	"github.com/GroupProject2-Kelompok4/BE/features/mentee"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
)

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

type searchMenteeResponse struct {
	No            uint   `json:"no"`
	Fullname      string `json:"name" form:"name"`
	ClassName     string `json:"class" form:"class"`
	Status        string `json:"status" form:"status"`
	EducationType string `json:"category" form:"category"`
	Gender        string `json:"gender" form:"gender"`
}

func searchMentee(m mentee.MenteeCore) searchMenteeResponse {
	return searchMenteeResponse{
		No:            m.No,
		Fullname:      m.Fullname,
		ClassName:     m.ClassName,
		Status:        m.Status,
		EducationType: m.EducationType,
		Gender:        m.Gender,
	}
}

type profileMenteeAndFeedbackResponse struct {
	MenteeID  string     `json:"mentee_id"`
	Users     string     `json:"users"`
	Feedbacks []Feedback `json:"feedbacks,omitempty"`
}

type Feedback struct {
	FeedbackId uint   `json:"feedback_id"`
	Status     string `json:"status"`
	Notes      string `json:"notes"`
	Proof      string `json:"proof"`
}

func profileMenteeAndFeedback(m mentee.MenteeCore) profileMenteeAndFeedbackResponse {
	feedbacks := make([]Feedback, len(m.Feedbacks))
	for i, feedback := range m.Feedbacks {
		feedbacks[i] = Feedback{
			FeedbackId: feedback.FeedbackId,
			Status:     feedback.Status,
			Notes:      feedback.Notes,
			Proof:      feedback.Proof,
		}
	}

	response := profileMenteeAndFeedbackResponse{
		MenteeID:  m.MenteeID,
		Users:     m.Users,
		Feedbacks: feedbacks,
	}

	return response
}

type profileMenteeLogResponse struct {
	MenteeID        string              `json:"mentee_id"`
	Class           string              `json:"class"`
	Fullname        string              `json:"full_name"`
	Nickname        string              `json:"nick_name"`
	Email           string              `json:"email"`
	Phone           string              `json:"phone"`
	CurrentAddress  string              `json:"current_address"`
	HomeAddress     string              `json:"home_address"`
	Telegram        string              `json:"telegram"`
	Gender          string              `json:"gender"`
	EducationType   string              `json:"education_type"`
	Major           string              `json:"major"`
	GraduateDate    helper.LocalTime    `json:"graduate_date"`
	Institution     string              `json:"institution"`
	EmergencyName   string              `json:"emergency_name"`
	EmergencyPhone  string              `json:"emergency_phone"`
	EmergencyStatus string              `json:"emergency_status"`
	Feedbacks       []FeedbackMenteeLog `json:"feedbacks,omitempty"`
}

type FeedbackMenteeLog struct {
	FeedbackId uint   `json:"feedback_id"`
	Users      string `json:"users"`
	Status     string `json:"status"`
	Notes      string `json:"notes"`
	Proof      string `json:"proof"`
}

func profileMenteeLog(m mentee.MenteeCore) profileMenteeLogResponse {
	feedbacks := make([]FeedbackMenteeLog, len(m.Feedbacks))
	for i, feedback := range m.Feedbacks {
		feedbacks[i] = FeedbackMenteeLog{
			FeedbackId: feedback.FeedbackId,
			Users:      feedback.Users,
			Status:     feedback.Status,
			Notes:      feedback.Notes,
			Proof:      feedback.Proof,
		}
	}

	resp := profileMenteeLogResponse{
		MenteeID:        m.MenteeID,
		Class:           m.ClassName,
		Fullname:        m.Fullname,
		Nickname:        m.Nickname,
		Email:           m.Email,
		Phone:           m.Phone,
		CurrentAddress:  m.CurrentAddress,
		HomeAddress:     m.HomeAddress,
		Telegram:        m.Telegram,
		Gender:          m.Gender,
		EducationType:   m.EducationType,
		Major:           m.Major,
		GraduateDate:    helper.LocalTime(m.GraduateDate),
		Institution:     m.Institution,
		EmergencyName:   m.EmergencyName,
		EmergencyPhone:  m.EmergencyPhone,
		EmergencyStatus: m.EmergencyStatus,
		Feedbacks:       feedbacks,
	}

	return resp
}
