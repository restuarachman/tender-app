package usecase

import (
	"log"
	"myapp/src/user"
	userDTO "myapp/src/user/dto"
	"myapp/src/user/entity"

	"github.com/go-playground/validator/v10"

	"myapp/pkg/middleware"
	"myapp/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo  user.UserRepository
	validator *validator.Validate
}

func NewUserUsecase(userRepo user.UserRepository, validator *validator.Validate) user.UserUsecase {
	return &userUsecase{
		userRepo:  userRepo,
		validator: validator,
	}
}

func (uu *userUsecase) Register(registerInfo userDTO.UserRegisterRequest) (userDTO.UserRegisterResponse, error) {
	if err := uu.validator.Struct(registerInfo); err != nil {
		return userDTO.UserRegisterResponse{}, err
	}

	_, err := uu.userRepo.FindByEmail(registerInfo.Email)
	if err.Error() != "record not found" {
		return userDTO.UserRegisterResponse{}, utils.ErrInternalServerError
	}

	if registerInfo.Password != registerInfo.PasswordConfirmation {
		return userDTO.UserRegisterResponse{}, utils.ErrPasswordDontMatch
	}

	userEntity := entity.User{
		Nickname:        registerInfo.Nickname,
		Email:           registerInfo.Email,
		Password:        hashAndSalt([]byte(registerInfo.Password)),
		ProfileImageUrl: registerInfo.ProfileImageUrl,
		Gender:          entity.Gender(registerInfo.Gender),
		Details:         registerInfo.Details,
	}

	_, err = uu.userRepo.Save(userEntity)
	if err != nil {
		return userDTO.UserRegisterResponse{}, err
	}

	token, err := middleware.JWTCreateToken(int(userEntity.ID), userEntity.Email, false)
	if err != nil {
		return userDTO.UserRegisterResponse{}, err
	}

	userRegisterDTO := userDTO.UserRegisterResponse{
		Nickname:        userEntity.Nickname,
		Email:           userEntity.Email,
		ProfileImageUrl: userEntity.ProfileImageUrl,
		Gender:          string(userEntity.Gender),
		Details:         userEntity.Details,
		Token:           token,
	}

	return userRegisterDTO, nil
}

func (uu *userUsecase) Login(loginInfo userDTO.UserLoginRequest) (userDTO.UserLoginResponse, error) {
	if err := uu.validator.Struct(loginInfo); err != nil {
		return userDTO.UserLoginResponse{}, err
	}

	userEntity, err := uu.userRepo.FindByEmail(loginInfo.Email)
	if err != nil {
		return userDTO.UserLoginResponse{}, utils.ErrInternalServerError
	}

	if !comparePasswords(userEntity.Password, []byte(loginInfo.Password)) {
		return userDTO.UserLoginResponse{}, utils.ErrBadParamInput
	}

	token, err := middleware.JWTCreateToken(int(userEntity.ID), userEntity.Email, userEntity.IsVerified)
	if err != nil {
		return userDTO.UserLoginResponse{}, err
	}

	userLoginDTO := userDTO.UserLoginResponse{
		Token: token,
	}

	return userLoginDTO, nil
}

func (uu *userUsecase) Update(updateInfo userDTO.UserUpdateRequest, userID uint) (userDTO.UserUpdateResponse, error) {
	if err := uu.validator.Struct(updateInfo); err != nil {
		return userDTO.UserUpdateResponse{}, err
	}

	userEntity, err := uu.userRepo.FindById(userID)
	if err != nil {
		return userDTO.UserUpdateResponse{}, utils.ErrInternalServerError
	}

	userEntity.Nickname = updateInfo.Nickname
	userEntity.ProfileImageUrl = updateInfo.ProfileImageUrl
	userEntity.Details = updateInfo.Details

	_, err = uu.userRepo.Update(userEntity, userID)
	if err != nil {
		return userDTO.UserUpdateResponse{}, err
	}

	userUpdateDTO := userDTO.UserUpdateResponse{
		Nickname:        userEntity.Nickname,
		ProfileImageUrl: userEntity.ProfileImageUrl,
		Details:         userEntity.Details,
	}

	return userUpdateDTO, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println("err", err)
	}

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println("err", err)
		return false
	}

	return true
}
