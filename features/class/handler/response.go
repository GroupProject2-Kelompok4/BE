package handler

import "github.com/GroupProject2-Kelompok4/BE/features/class"

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
