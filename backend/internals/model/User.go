package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"userName"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type UserCredentials struct {
	ID       uint
	Password string
	Email    string
}

type UserInfo struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

//Api structs

type UserSignupPayload struct {
	UserName        string `json:"userName" validate:"required"`
	Email           string `json:"email" validate:"required,min=3"`
	Password        string `json:"password" validate:"required min=6"`
	ConfirmPassword string `json:"confirmPassword" validate:"required eqfield=Password"`
}

type GetUserByEmailPayload struct {
	Email string `json:"email" validate:"required,min=3"`
}

type UserLoginPayload struct {
	Email    string `json:"email" validate:"required,min=3"`
	Password string `json:"password" validate:"required min=6"`
}

type UserLoginResponse struct {
	Message string `json:"message"`
	Jwt     string `json:"jwt"`
}

type UserProfileResponse struct {
	UserInfo
	Message string `json:"message"`
}
