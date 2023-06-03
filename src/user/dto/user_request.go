package dto

type UserRegisterRequest struct {
	Nickname             string `json:"nickname" validate:"required"`
	Email                string `json:"email" validate:"required"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
	ProfileImageUrl      string `json:"profile_image_url" validate:"required"`
	Gender               string `json:"gender" validate:"required"`
	Details              string `json:"details"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Nickname        string `json:"nickname"`
	ProfileImageUrl string `json:"profile_image_url"`
	Details         string `json:"details"`
}
