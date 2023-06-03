package main

import (
	"log"
	"net/http"

	_config "myapp/config"
	_driver "myapp/src/driver"
	_userHandler "myapp/src/user/handler"
	_userRepo "myapp/src/user/repository/mysql"
	_userUsecase "myapp/src/user/usecase"

	_interactionHandler "myapp/src/interaction/handler"
	_interactionRepo "myapp/src/interaction/repository/mysql"
	_interactionUsecase "myapp/src/interaction/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config, err := _config.LoadConfig(".")
	if err != nil {
		log.Fatal("err", err)
	}
	_driver.ConnectDB(
		config.DBConn,
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPass,
		config.DBName,
	)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	v := validator.New()

	// setup middleware
	JWTSecret, err := _config.LoadJWTSecret(".")
	if err != nil {
		log.Fatal("err", err)
	}

	// setup repository
	userRepo := _userRepo.NewMysqlUserRepository(_driver.DB)
	interactionRepo := _interactionRepo.NewMysqlInteractionRepository(_driver.DB)

	// setup usecase
	userUsecase := _userUsecase.NewUserUsecase(userRepo, v)
	interactionUsecase := _interactionUsecase.NewInteractionUsecase(interactionRepo, userRepo, v)

	// setup route
	_userHandler.NewUserHandler(e, userUsecase, JWTSecret.Secret)
	_interactionHandler.NewInteractionHandler(e, interactionUsecase, JWTSecret.Secret)

	// setup middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	log.Fatal(e.Start(":" + config.ServerPort))
}
