package routes

import (
	"github.com/labstack/echo/v4"
	"go-user_api_example/controllers"
)

func UserRoute(e *echo.Echo) {
	e.POST("/user", controllers.CreateUser)
}
