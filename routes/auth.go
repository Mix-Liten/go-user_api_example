package routes

import (
	"github.com/labstack/echo/v4"
	"go-user_api_example/modules/common/handler"
)

func AuthRoute(e *echo.Echo) {

	api := e.Group("/api")
	api.POST("/auth", handler.CreateToken)
}
