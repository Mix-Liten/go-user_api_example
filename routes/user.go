package routes

import (
	"github.com/labstack/echo/v4"
	"go-user_api_example/configs/database"
	"go-user_api_example/modules/user/handler"
	"go-user_api_example/modules/user/repository"
)

func UserRoute(e *echo.Echo) {
	ur := repository.NewUserRepositoryMongo(database.GetDB(), "users")
	uh := handler.NewUserHandler(ur)

	e.POST("/user", uh.CreateUser)
}
