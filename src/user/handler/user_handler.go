package handler

import (
	"myapp/pkg/response"
	"myapp/src/user"
	userDTO "myapp/src/user/dto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_midl "myapp/pkg/middleware"
)

type UserHandler struct {
	userUsecase user.UserUsecase
	JWTSecret   string
}

func NewUserHandler(e *echo.Echo, us user.UserUsecase, JWTSecret string) {
	handler := &UserHandler{
		userUsecase: us,
		JWTSecret:   JWTSecret,
	}

	e.POST("/api/v1/users/register", handler.Register)
	e.POST("/api/v1/users/login", handler.Login)
	e.PUT("/api/v1/users/update", handler.Update, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/users/upgrade", handler.UpgradeAccount, middleware.JWT([]byte(JWTSecret)))
}

func (u *UserHandler) Login(c echo.Context) error {
	loginInfo := userDTO.UserLoginRequest{}

	c.Bind(&loginInfo)

	userResponse, err := u.userUsecase.Login(loginInfo)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, userResponse)
}

func (u *UserHandler) Register(c echo.Context) error {
	user := userDTO.UserRegisterRequest{}
	c.Bind(&user)

	userResponse, err := u.userUsecase.Register(user)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, userResponse)
}

func (u *UserHandler) Update(c echo.Context) error {
	user := userDTO.UserUpdateRequest{}
	c.Bind(&user)

	userID, _, _ := _midl.ExtractTokenUser(c)

	userReponse, err := u.userUsecase.Update(user, uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, userReponse)
}

func (u *UserHandler) UpgradeAccount(c echo.Context) error {
	userID, _, _ := _midl.ExtractTokenUser(c)

	message, err := u.userUsecase.UpgradeAccount(uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, message)
}
