package model

type ResetPasswordMailData struct {
	Email    string `json:"email"`
	ResetURL string `json:"reset_url"`
	Year     int    `json:"year"`
	Token    string `json:"token"`
}

type ChangePasswordByEmailPayload struct {
	Password string `json:"password"`
}
