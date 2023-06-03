package interaction

import (
	interactionDTO "myapp/src/interaction/dto"
	userDTO "myapp/src/user/dto"
)

type InteractionUsecase interface {
	Swipe(userGivenId uint, interactionDTO interactionDTO.InteractionWithUser) (interactionDTO.InteractionUser, error)
	ShowRandomPeople(userId uint) (userDTO.UserShowResponse, error)
}
