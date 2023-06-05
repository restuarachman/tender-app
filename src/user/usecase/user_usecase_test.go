package usecase

import (
	"errors"
	"myapp/src/user/dto"
	"myapp/src/user/entity"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"myapp/src/user/mocks"
)

var (
	userEntity = entity.User{
		Model: gorm.Model{
			ID:        uint(1),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Nickname:        "test",
		Email:           "test@gmail.com",
		Password:        "test123",
		ProfileImageUrl: "https://image.png",
		Gender:          entity.Gender("Female"),
		Popularity:      0,
		IsVerified:      false,
		Details:         "test",
	}
)

func TestRegister(t *testing.T) {

	mockUserRepo := mocks.NewUserRepository(t)

	registerInfo := dto.UserRegisterRequest{
		Nickname:             "test",
		Email:                "test@gmail.com",
		Password:             "test123",
		PasswordConfirmation: "test123",
		ProfileImageUrl:      "https://image.png",
		Gender:               "Female",
		Details:              "test",
	}

	userEntity2 := entity.User{
		Nickname:        registerInfo.Nickname,
		Email:           registerInfo.Email,
		Password:        "test123",
		ProfileImageUrl: registerInfo.ProfileImageUrl,
		Gender:          entity.Gender(registerInfo.Gender),
		Details:         registerInfo.Details,
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", userEntity.Email).Return(entity.User{}, errors.New("record not found")).Once()
		mockUserRepo.On("Save", userEntity2).Return(userEntity, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, validator.New())
		res, err := testUserUsecase.Register(registerInfo)

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})
}
