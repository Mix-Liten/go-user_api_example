package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-user_api_example/configs/database"
	ch "go-user_api_example/modules/common/handler"
	"go-user_api_example/modules/user/handler"
	"go-user_api_example/modules/user/repository"
	"os"
)

func UserRoute(e *echo.Echo) {
	ur := repository.NewUserRepositoryMongo(database.GetDB(), "users")
	uh := handler.NewUserHandler(ur)

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{Claims: &ch.TokenClaims{}, SigningKey: []byte(os.Getenv("JWT_SECRET"))}))
	api.POST("/user", uh.CreateUser)
	api.GET("/user/:userID", uh.GetUser)
	api.PUT("/user/:userID", uh.EditUser)
	api.DELETE("/user/:userID", uh.DeleteUser)
	api.GET("/users", uh.GetAllUser)
}
