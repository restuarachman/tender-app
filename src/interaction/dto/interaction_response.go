package dto

import userDTO "myapp/src/user/dto"

type InteractionUser struct {
	userDTO.UserShowResponse
	InteractionType string `json:"interaction_type"`
	CreatedAt       string `json:"created_at"`
}
