package handler

import (
	_midl "myapp/pkg/middleware"
	"myapp/pkg/response"
	"myapp/src/interaction"
	"myapp/src/interaction/dto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type InteractionHandler struct {
	interactionUsecase interaction.InteractionUsecase
	JWTSecret          string
}

func NewInteractionHandler(e *echo.Echo, iu interaction.InteractionUsecase, JWTSecret string) {
	handler := &InteractionHandler{interactionUsecase: iu, JWTSecret: JWTSecret}
	e.POST("api/v1/interactions/swipe", handler.Swipe, middleware.JWT([]byte(JWTSecret)))
	e.GET("api/v1/interactions/show-people", handler.ShowRandomPeople, middleware.JWT([]byte(JWTSecret)))
}

func (ih *InteractionHandler) Swipe(c echo.Context) error {
	userGivenId, _, _ := _midl.ExtractTokenUser(c)

	interactUser := dto.InteractionWithUser{}
	c.Bind(&interactUser)

	interactionResponse, err := ih.interactionUsecase.Swipe(uint(userGivenId), interactUser)
	if err != nil {
		return err
	}

	return response.SuccessResponse(c, interactionResponse)
}

func (ih *InteractionHandler) ShowRandomPeople(c echo.Context) error {
	userGivenId, _, _ := _midl.ExtractTokenUser(c)

	userResponse, err := ih.interactionUsecase.ShowRandomPeople(uint(userGivenId))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, userResponse)
}
