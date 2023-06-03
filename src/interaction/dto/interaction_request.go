package dto

type InteractionWithUser struct {
	UserId          uint   `json:"user_id"`
	InteractionType string `json:"interaction_type"`
}
