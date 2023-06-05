package user

import (
	userDTO "myapp/src/user/dto"
)

type UserUsecase interface {
	Register(registerInfo userDTO.UserRegisterRequest) (userDTO.UserRegisterResponse, error)
	Login(loginInfo userDTO.UserLoginRequest) (userDTO.UserLoginResponse, error)
	Update(updateInfo userDTO.UserUpdateRequest, userID uint) (userDTO.UserUpdateResponse, error)
	UpgradeAccount(userID uint) (string, error)
}
