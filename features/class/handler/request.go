package handler

import (
	"time"

	"github.com/GroupProject2-Kelompok4/BE/features/class"
	"github.com/GroupProject2-Kelompok4/BE/utils"
)

var log = utils.Log()

type RegisterClassRequest struct {
	Name         string `json:"name" form:"name"`
	UserID       string `json:"pic" form:"pic"`
	StartDate    string `json:"start_date" form:"start_date"`
	GraduateDate string `json:"graduate_date" form:"graduate_date"`
}

func RequestToCore(data interface{}) class.ClassCore {
	res := class.ClassCore{}
	switch v := data.(type) {
	case RegisterClassRequest:
		res.Name = v.Name
		res.UserID = v.UserID
		startDate, err := time.ParseInLocation("2006-01-02", v.StartDate, time.Local)
		if err != nil {
			log.Error("error while parsing string to time format")
			return class.ClassCore{}
		}
		res.StartDate = startDate

		graduateDate, err := time.ParseInLocation("2006-01-02", v.GraduateDate, time.Local)
		if err != nil {
			log.Error("error while parsing string to time format")
			return class.ClassCore{}
		}
		res.GraduateDate = graduateDate

	default:
		return class.ClassCore{}
	}
	return res
}
