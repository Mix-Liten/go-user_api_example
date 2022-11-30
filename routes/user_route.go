package routes

import (
	"github.com/labstack/echo/v4"
	"go-user_api_example/controllers"
)

func UserRoute(e *echo.Echo) {
	uc := controllers.UserController{}
	
	e.POST("/user", uc.CreateUser)
}
