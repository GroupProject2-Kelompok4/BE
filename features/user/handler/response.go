package handler

import "github.com/GroupProject2-Kelompok4/BE/features/user"

type loginResponse struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
	Role   string `json:"role"`
}

type registerResponse struct {
	UserID   string `json:"user_id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func register(u user.UserCore) registerResponse {
	return registerResponse{
		UserID:   u.UserID,
		Fullname: u.Fullname,
		Email:    u.Email,
		Role:     u.Role,
	}
}
