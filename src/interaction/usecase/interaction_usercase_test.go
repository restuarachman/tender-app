package usecase

import (
	_interactionDto "myapp/src/interaction/dto"
	_interactionEntity "myapp/src/interaction/entity"
	_interactionMock "myapp/src/interaction/mocks"
	_userEntity "myapp/src/user/entity"
	_userMock "myapp/src/user/mocks"
	"testing"
	"time"

	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	mockUserEntity = _userEntity.User{
		Model: gorm.Model{
			ID:        uint(1),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Nickname:        "test",
		Email:           "test@gmail.com",
		Password:        "test123",
		ProfileImageUrl: "https://image.png",
		Gender:          _userEntity.Gender("Female"),
		Popularity:      0,
		IsVerified:      false,
		Details:         "test",
	}

	mockUserEntity2 = _userEntity.User{
		Model: gorm.Model{
			ID:        uint(2),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Nickname:        "test1",
		Email:           "test1@gmail.com",
		Password:        "test123",
		ProfileImageUrl: "https://image.png",
		Gender:          _userEntity.Gender("Female"),
		Popularity:      0,
		IsVerified:      false,
		Details:         "test",
	}

	mockUserEntity3 = _userEntity.User{
		Model: gorm.Model{
			ID:        uint(3),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Nickname:        "test2",
		Email:           "test2@gmail.com",
		Password:        "test123",
		ProfileImageUrl: "https://image.png",
		Gender:          _userEntity.Gender("Female"),
		Popularity:      0,
		IsVerified:      false,
		Details:         "test",
	}

	mockUserEntity4 = _userEntity.User{
		Model: gorm.Model{
			ID:        uint(4),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Nickname:        "test4",
		Email:           "test4@gmail.com",
		Password:        "test123",
		ProfileImageUrl: "https://image.png",
		Gender:          _userEntity.Gender("Female"),
		Popularity:      0,
		IsVerified:      false,
		Details:         "test",
	}

	userEntityArr = []_userEntity.User{mockUserEntity2, mockUserEntity3}

	mockIntEntity = _interactionEntity.Interaction{
		UserGivenId:     uint(1),
		UserReceivedId:  uint(2),
		InteractionType: _interactionEntity.InteractionType("like"),
		CreatedAt:       time.Now(),
	}

	mockIntEntity2 = _interactionEntity.Interaction{
		UserGivenId:     uint(1),
		UserReceivedId:  uint(3),
		InteractionType: _interactionEntity.InteractionType("like"),
		CreatedAt:       time.Now(),
	}

	mockIntEntityArr             = []_interactionEntity.Interaction{mockIntEntity, mockIntEntity2}
	mockInteEntityArrMoreThanTen = []_interactionEntity.Interaction{mockIntEntity, mockIntEntity2, mockIntEntity, mockIntEntity2, mockIntEntity, mockIntEntity2, mockIntEntity, mockIntEntity2, mockIntEntity, mockIntEntity2, mockIntEntity, mockIntEntity2}
)

func TestShowRandomPeople(t *testing.T) {
	mockInteractionRepo := _interactionMock.NewInteractionRepository(t)
	mockUserRepo := _userMock.NewUserRepository(t)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindById", uint(1)).Return(mockUserEntity, nil).Once()
		mockInteractionRepo.On("FindCurrentInteraction", uint(1)).Return(mockIntEntityArr, nil).Once()
		mockInteractionRepo.On("FindRandomPeople", uint(1)).Return(mockUserEntity4, nil).Once()

		testInteractionUsecase := NewInteractionUsecase(mockInteractionRepo, mockUserRepo, nil)
		res, err := testInteractionUsecase.ShowRandomPeople(uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("user-error", func(t *testing.T) {
		mockUserRepo.On("FindById", uint(1)).Return(_userEntity.User{}, errors.New("error")).Once()

		testInteractionUsecase := NewInteractionUsecase(mockInteractionRepo, mockUserRepo, nil)
		res, err := testInteractionUsecase.ShowRandomPeople(uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("current-interaction-error", func(t *testing.T) {
		mockUserRepo.On("FindById", uint(1)).Return(mockUserEntity, nil).Once()
		mockInteractionRepo.On("FindCurrentInteraction", uint(1)).Return([]_interactionEntity.Interaction{}, errors.New("error")).Once()

		testInteractionUsecase := NewInteractionUsecase(mockInteractionRepo, mockUserRepo, nil)
		res, err := testInteractionUsecase.ShowRandomPeople(uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("action-limit-error", func(t *testing.T) {
		mockUserRepo.On("FindById", uint(1)).Return(mockUserEntity, nil).Once()
		mockInteractionRepo.On("FindCurrentInteraction", uint(1)).Return(mockInteEntityArrMoreThanTen, nil).Once()

		testInteractionUsecase := NewInteractionUsecase(mockInteractionRepo, mockUserRepo, nil)
		res, err := testInteractionUsecase.ShowRandomPeople(uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("find-people-error", func(t *testing.T) {
		mockUserRepo.On("FindById", uint(1)).Return(mockUserEntity, nil).Once()
		mockInteractionRepo.On("FindCurrentInteraction", uint(1)).Return(mockIntEntityArr, nil).Once()
		mockInteractionRepo.On("FindRandomPeople", uint(1)).Return(_userEntity.User{}, errors.New("error")).Once()

		testInteractionUsecase := NewInteractionUsecase(mockInteractionRepo, mockUserRepo, nil)
		res, err := testInteractionUsecase.ShowRandomPeople(uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestSwipe(t *testing.T) {
	mockInteractionRepo := _interactionMock.NewInteractionRepository(t)
	mockUserRepo := _userMock.NewUserRepository(t)

	mockInteractionEntity := _interactionEntity.Interaction{
		UserGivenId:     uint(1),
		UserReceivedId:  uint(4),
		InteractionType: _interactionEntity.Like,
	}

	mockInteractionEntityReturn := _interactionEntity.Interaction{
		UserGivenId:     uint(1),
		UserReceivedId:  uint(4),
		InteractionType: _interactionEntity.Like,
		CreatedAt:       time.Now(),
	}

	mockInteractionDto := _interactionDto.InteractionWithUser{
		UserId:          uint(4),
		InteractionType: "like",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindById", uint(1)).Return(mockUserEntity, nil).Once()
		mockInteractionRepo.On("FindCurrentInteraction", uint(1)).Return(mockIntEntityArr, nil).Once()
		mockInteractionRepo.On("Save", mockInteractionEntity).Return(mockInteractionEntityReturn, nil).Once()

		testInteractionUsecase := NewInteractionUsecase(mockInteractionRepo, mockUserRepo, validator.New())
		res, err := testInteractionUsecase.Swipe(uint(1), mockInteractionDto)

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("user-error", func(t *testing.T) {
		mockUserRepo.On("FindById", uint(1)).Return(_userEntity.User{}, errors.New("error")).Once()

		testInteractionUsecase := NewInteractionUsecase(mockInteractionRepo, mockUserRepo, nil)
		res, err := testInteractionUsecase.Swipe(uint(1), mockInteractionDto)

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("current-interaction-error", func(t *testing.T) {
		mockUserRepo.On("FindById", uint(1)).Return(mockUserEntity, nil).Once()
		mockInteractionRepo.On("FindCurrentInteraction", uint(1)).Return([]_interactionEntity.Interaction{}, errors.New("error")).Once()

		testInteractionUsecase := NewInteractionUsecase(mockInteractionRepo, mockUserRepo, nil)
		res, err := testInteractionUsecase.Swipe(uint(1), mockInteractionDto)

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("action-limit-error", func(t *testing.T) {
		mockUserRepo.On("FindById", uint(1)).Return(mockUserEntity, nil).Once()
		mockInteractionRepo.On("FindCurrentInteraction", uint(1)).Return(mockInteEntityArrMoreThanTen, nil).Once()

		testInteractionUsecase := NewInteractionUsecase(mockInteractionRepo, mockUserRepo, nil)
		res, err := testInteractionUsecase.Swipe(uint(1), mockInteractionDto)

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}
