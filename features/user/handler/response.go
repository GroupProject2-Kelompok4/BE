package handler

import (
	"github.com/GroupProject2-Kelompok4/BE/features/user"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
)

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
	Team     string `json:"team"`
}

func register(u user.UserCore) registerResponse {
	return registerResponse{
		UserID:   u.UserID,
		Fullname: u.Fullname,
		Email:    u.Email,
		Role:     u.Role,
		Team:     u.Team,
	}
}

type searchUserResponse struct {
	No        uint             `json:"no"`
	UserID    string           `json:"user_id"`
	Fullname  string           `json:"fullname"`
	Email     string           `json:"email"`
	Team      string           `json:"team"`
	Role      string           `json:"role"`
	Status    string           `json:"status"`
	CreatedAt helper.LocalTime `json:"created_at"`
	UpdatedAt helper.LocalTime `json:"updated_at"`
}

func searchUser(u user.UserCore) searchUserResponse {
	return searchUserResponse{
		No:        u.No,
		UserID:    u.UserID,
		Fullname:  u.Fullname,
		Email:     u.Email,
		Team:      u.Team,
		Role:      u.Role,
		Status:    u.Status,
		CreatedAt: helper.LocalTime(u.CreatedAt),
		UpdatedAt: helper.LocalTime(u.UpdatedAt),
	}
}

type profileResponse struct {
	UserID    string           `json:"user_id"`
	Fullname  string           `json:"fullname"`
	Email     string           `json:"email"`
	Team      string           `json:"team"`
	Role      string           `json:"role"`
	Status    string           `json:"status"`
	CreatedAt helper.LocalTime `json:"created_at"`
	UpdatedAt helper.LocalTime `json:"updated_at"`
}

func profileUser(u user.UserCore) profileResponse {
	return profileResponse{
		UserID:    u.UserID,
		Fullname:  u.Fullname,
		Email:     u.Email,
		Team:      u.Team,
		Role:      u.Role,
		Status:    u.Status,
		CreatedAt: helper.LocalTime(u.CreatedAt),
		UpdatedAt: helper.LocalTime(u.UpdatedAt),
	}
}

type updateResponse struct {
	Fullname  string           `json:"fullname"`
	Email     string           `json:"email"`
	Team      string           `json:"team"`
	Role      string           `json:"role"`
	Status    string           `json:"status"`
	CreatedAt helper.LocalTime `json:"created_at"`
	UpdatedAt helper.LocalTime `json:"updated_at"`
}

func updateUserProfile(u user.UserCore) updateResponse {
	return updateResponse{
		Fullname:  u.Fullname,
		Email:     u.Email,
		Team:      u.Team,
		Role:      u.Role,
		Status:    u.Status,
		CreatedAt: helper.LocalTime(u.CreatedAt),
		UpdatedAt: helper.LocalTime(u.UpdatedAt),
	}
}
