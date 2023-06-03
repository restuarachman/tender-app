package dto

type UserRegisterRequest struct {
	FullName             string `json:"full_name" validate:"required"`
	Username             string `json:"username" validate:"required"`
	Email                string `json:"email" validate:"required"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

type UserLoginRequest struct {
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
}
