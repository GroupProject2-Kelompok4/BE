package handler

import "github.com/GroupProject2-Kelompok4/BE/features/user"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Team     string `json:"team" form:"team"`
	Status   string `json:"status" form:"status"`
	Role     string `json:"role" form:"role"`
}

type UpdateProfileRequest struct {
	Fullname *string `json:"fullname" form:"fullname"`
	Email    *string `json:"email" form:"email"`
	Password *string `json:"password" form:"password"`
	Team     *string `json:"team" form:"team"`
	Status   *string `json:"status" form:"status"`
	Role     *string `json:"role" form:"role"`
}

func RequestToCore(data interface{}) user.UserCore {
	res := user.UserCore{}
	switch v := data.(type) {
	case LoginRequest:
		res.Email = v.Email
		res.Password = v.Password
	case RegisterRequest:
		res.Fullname = v.Fullname
		res.Email = v.Email
		res.Password = v.Password
		res.Team = v.Team
		res.Status = v.Status
		res.Role = v.Role
	case *UpdateProfileRequest:
		if v.Fullname != nil {
			res.Fullname = *v.Fullname
		}
		if v.Email != nil {
			res.Email = *v.Email
		}
		if v.Password != nil {
			res.Password = *v.Password
		}
		if v.Team != nil {
			res.Team = *v.Team
		}
		if v.Status != nil {
			res.Status = *v.Status
		}
		if v.Role != nil {
			res.Role = *v.Role
		}
	default:
		return user.UserCore{}
	}
	return res
}
