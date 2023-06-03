package user

import (
	userDTO "myapp/src/user/dto"
)

type UserUsecase interface {
	Register(registerInfo userDTO.UserRegisterRequest) (userDTO.UserRegisterResponse, error)
	Login(loginInfo userDTO.UserLoginRequest) (userDTO.UserLoginResponse, error)
}
