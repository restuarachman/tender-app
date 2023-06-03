package interaction

import (
	interactionEntity "myapp/src/interaction/entity"
	userEntity "myapp/src/user/entity"
)

type InteractionRepository interface {
	FindByUserGivenId(id uint) ([]interactionEntity.Interaction, error)
	FindByUserReceivedId(id uint) ([]interactionEntity.Interaction, error)
	Save(interaction interactionEntity.Interaction) (interactionEntity.Interaction, error)
	FindCurrentInteraction(userId uint) ([]interactionEntity.Interaction, error)
	FindRandomPeople(userId uint) (userEntity.User, error)
}
