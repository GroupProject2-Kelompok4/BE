package handler

import (
	"github.com/GroupProject2-Kelompok4/BE/features/class"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
)

type registerClassResponse struct {
	ClassID string `json:"class_id"`
	Name    string `json:"name"`
	PIC     string `json:"pic"`
}

func registerClass(c class.ClassCore, pic string) registerClassResponse {
	return registerClassResponse{
		ClassID: c.ClassID,
		Name:    c.Name,
		PIC:     pic,
	}
}

type listClassesResponse struct {
	No           uint             `json:"no"`
	ClassID      string           `json:"class_id"`
	Name         string           `json:"name"`
	PIC          string           `json:"pic"`
	StartDate    helper.LocalTime `json:"start_date"`
	GraduateDate helper.LocalTime `json:"graduate_date"`
}

func listClasses(c class.ClassCore) listClassesResponse {
	return listClassesResponse{
		No:           c.No,
		ClassID:      c.ClassID,
		Name:         c.Name,
		PIC:          c.PIC,
		StartDate:    helper.LocalTime(c.StartDate),
		GraduateDate: helper.LocalTime(c.GraduateDate),
	}
}

type getClassResponse struct {
	ClassID      string           `json:"class_id"`
	Name         string           `json:"name"`
	PIC          string           `json:"pic"`
	StartDate    helper.LocalTime `json:"start_date"`
	GraduateDate helper.LocalTime `json:"graduate_date"`
}

func getClass(c class.ClassCore) getClassResponse {
	return getClassResponse{
		ClassID:      c.ClassID,
		Name:         c.Name,
		PIC:          c.PIC,
		StartDate:    helper.LocalTime(c.StartDate),
		GraduateDate: helper.LocalTime(c.GraduateDate),
	}
}
