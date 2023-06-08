package entities

type RegisterRequest struct {
	Username       string `json:"username"`
	Fullname       string `json:"fullname"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ForgotPasswordRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}
