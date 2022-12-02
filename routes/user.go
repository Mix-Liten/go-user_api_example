package routes

import (
	"github.com/labstack/echo/v4"
	"go-user_api_example/modules/user/handler"
)

func UserRoute(e *echo.Echo) {
	uh := handler.UserHandler{}

	e.POST("/user", uh.CreateUser)
}
