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

	_, err := uu.userRepo.FindByUsernameOrEmail(registerInfo.Username, registerInfo.Email)
	if err.Error() != "record not found" {
		return userDTO.UserRegisterResponse{}, utils.ErrInternalServerError
	}

	if registerInfo.Password != registerInfo.PasswordConfirmation {
		return userDTO.UserRegisterResponse{}, utils.ErrPasswordDontMatch
	}

	userEntity := entity.User{
		Email:              registerInfo.Email,
		Username:           registerInfo.Username,
		Password:           hashAndSalt([]byte(registerInfo.Password)),
		ProfileImageUrl:    "https://profile.png",
		BackgroundImageUrl: "https://bg-profile.png",
		Name:               registerInfo.FullName,
		ValidEmail:         false,
	}

	_, err = uu.userRepo.Save(userEntity)
	if err != nil {
		return userDTO.UserRegisterResponse{}, err
	}

	token, err := middleware.JWTCreateToken(int(userEntity.ID), userEntity.Username, userEntity.Email)
	if err != nil {
		return userDTO.UserRegisterResponse{}, err
	}

	userRegisterDTO := userDTO.UserRegisterResponse{
		Email:    userEntity.Email,
		Username: userEntity.Username,
		FullName: userEntity.Name,
		Token:    token,
	}

	return userRegisterDTO, nil
}

func (uu *userUsecase) Login(loginInfo userDTO.UserLoginRequest) (userDTO.UserLoginResponse, error) {
	if err := uu.validator.Struct(loginInfo); err != nil {
		return userDTO.UserLoginResponse{}, err
	}

	userEntity, err := uu.userRepo.FindByUsernameOrEmail(loginInfo.User, loginInfo.User)
	if err != nil {
		return userDTO.UserLoginResponse{}, utils.ErrInternalServerError
	}

	if !comparePasswords(userEntity.Password, []byte(loginInfo.Password)) {
		return userDTO.UserLoginResponse{}, utils.ErrBadParamInput
	}

	token, err := middleware.JWTCreateToken(int(userEntity.ID), userEntity.Username, userEntity.Email)
	if err != nil {
		return userDTO.UserLoginResponse{}, err
	}

	userLoginDTO := userDTO.UserLoginResponse{
		Email:    userEntity.Email,
		Username: userEntity.Username,
		Token:    token,
	}

	return userLoginDTO, nil
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println("err", err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println("err", err)
		return false
	}

	return true
}
