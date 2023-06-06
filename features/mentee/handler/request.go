package handler

import (
	"time"

	"github.com/GroupProject2-Kelompok4/BE/features/mentee"
	"github.com/GroupProject2-Kelompok4/BE/utils"
)

var log = utils.Log()

type RegisterMenteeRequest struct {
	Fullname        string `json:"full_name" form:"full_name"`
	Nickname        string `json:"nickname" form:"nickname"`
	Phone           string `json:"phone" form:"phone"`
	Email           string `json:"email" form:"email"`
	CurrentAddress  string `json:"current_address" form:"current_address"`
	HomeAddress     string `json:"home_address" form:"home_address"`
	Telegram        string `json:"telegram" form:"telegram"`
	Gender          string `json:"gender" form:"gender"`
	EducationType   string `json:"education_type" form:"education_type"`
	Major           string `json:"major" form:"major"`
	GraduateDate    string `json:"graduate_date" form:"graduate_date"`
	Institution     string `json:"institution" form:"institution"`
	EmergencyName   string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status" form:"emergency_status"`
	Status          string `json:"status" form:"status"`
	ClassID         string `json:"class_id" form:"class_id"`
}

func RequestToCore(data interface{}) mentee.MenteeCore {
	res := mentee.MenteeCore{}
	switch v := data.(type) {
	case RegisterMenteeRequest:
		res.Fullname = v.Fullname
		res.Nickname = v.Nickname
		res.Phone = v.Phone
		res.Email = v.Email
		res.CurrentAddress = v.CurrentAddress
		res.HomeAddress = v.HomeAddress
		res.Telegram = v.Telegram
		res.Gender = v.Gender
		res.EducationType = v.EducationType
		res.Major = v.Major
		graduateDate, err := time.ParseInLocation("2006-01-02", v.GraduateDate, time.Local)
		if err != nil {
			log.Error("error while parsing string to time format")
			return mentee.MenteeCore{}
		}
		res.GraduateDate = graduateDate
		res.Institution = v.Institution
		res.EmergencyName = v.EmergencyName
		res.EmergencyPhone = v.EmergencyPhone
		res.EmergencyStatus = v.EmergencyStatus
		res.Status = v.Status
		res.ClassID = v.ClassID
	default:
		return mentee.MenteeCore{}
	}
	return res
}
