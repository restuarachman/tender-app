package usecase

import (
	"myapp/pkg/utils"
	"myapp/src/interaction"
	interactionDTO "myapp/src/interaction/dto"
	interactionEntity "myapp/src/interaction/entity"
	"myapp/src/user"
	userDTO "myapp/src/user/dto"

	"github.com/go-playground/validator/v10"
)

type interactionUsecase struct {
	interactionRepo interaction.InteractionRepository
	userRepo        user.UserRepository
	validator       *validator.Validate
}

func NewInteractionUsecase(interactionRepo interaction.InteractionRepository, userRepo user.UserRepository, validator *validator.Validate) interaction.InteractionUsecase {
	return &interactionUsecase{interactionRepo, userRepo, validator}
}

func (iu *interactionUsecase) Swipe(userGivenId uint, interactiondto interactionDTO.InteractionWithUser) (interactionDTO.InteractionUser, error) {
	userGiven, err := iu.userRepo.FindById(userGivenId)
	if err != nil {
		return interactionDTO.InteractionUser{}, err
	}

	interactions, err := iu.interactionRepo.FindCurrentInteraction(userGiven.ID)
	if err != nil {
		return interactionDTO.InteractionUser{}, err
	}

	// Limit the number of interactions per day if the user is not verified
	interactionCount := len(interactions)
	if interactionCount > 10 && !userGiven.IsVerified {
		return interactionDTO.InteractionUser{}, utils.ErrBadParamInput
	}

	interaction := interactionEntity.Interaction{
		UserGivenId:     userGivenId,
		UserReceivedId:  interactiondto.UserId,
		InteractionType: interactionEntity.InteractionType(interactiondto.InteractionType),
	}

	err = iu.validator.Struct(interaction)
	if err != nil {
		return interactionDTO.InteractionUser{}, err
	}

	interaction, err = iu.interactionRepo.Save(interaction)

	userResponse := userDTO.UserShowResponse{
		ID:              userGiven.ID,
		Nickname:        userGiven.Nickname,
		Email:           userGiven.Email,
		ProfileImageUrl: userGiven.ProfileImageUrl,
		Gender:          string(userGiven.Gender),
		Details:         userGiven.Details,
		IsVerified:      userGiven.IsVerified,
	}

	interactionResponse := interactionDTO.InteractionUser{
		UserShowResponse: userResponse,
		InteractionType:  string(interaction.InteractionType),
	}

	return interactionResponse, err
}

func (iu *interactionUsecase) ShowRandomPeople(userId uint) (userDTO.UserShowResponse, error) {
	userGiven, err := iu.userRepo.FindById(userId)
	if err != nil {
		return userDTO.UserShowResponse{}, err
	}

	interactions, err := iu.interactionRepo.FindCurrentInteraction(userGiven.ID)
	if err != nil {
		return userDTO.UserShowResponse{}, err
	}

	// Limit the number of interactions per day if the user is not verified
	interactionCount := len(interactions)
	if interactionCount > 10 && !userGiven.IsVerified {
		return userDTO.UserShowResponse{}, utils.ErrBadParamInput
	}

	randomUser, err := iu.interactionRepo.FindRandomPeople(userId)
	if err != nil {
		return userDTO.UserShowResponse{}, err
	}

	userResponse := userDTO.UserShowResponse{
		ID:              randomUser.ID,
		Nickname:        randomUser.Nickname,
		Email:           randomUser.Email,
		ProfileImageUrl: randomUser.ProfileImageUrl,
		Gender:          string(randomUser.Gender),
		Details:         randomUser.Details,
		IsVerified:      randomUser.IsVerified,
	}

	return userResponse, nil
}
