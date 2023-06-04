package handler

import "github.com/GroupProject2-Kelompok4/BE/features/user"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(data interface{}) user.UserCore {
	res := user.UserCore{}
	switch v := data.(type) {
	case LoginRequest:
		res.Email = v.Email
		res.Password = v.Password
	default:
		return user.UserCore{}
	}
	return res
}
