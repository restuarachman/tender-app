package dto

type UserRegisterResponse struct {
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	ProfileImageUrl string `json:"profile_image_url"`
	Gender          string `json:"gender"`
	Details         string `json:"details"`
	Token           string `json:"token"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserUpdateResponse struct {
	Nickname        string `json:"nickname"`
	ProfileImageUrl string `json:"profile_image_url"`
	Details         string `json:"details"`
}
