package handler

type LoginResponse struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
	Role   string `json:"role"`
}
