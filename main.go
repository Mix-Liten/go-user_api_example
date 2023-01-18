package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-user_api_example/configs"
	"go-user_api_example/configs/database"
	"go-user_api_example/helpers"
	"go-user_api_example/routes"
	"net/http"
)

func init() {
	helpers.VerifyEnv()
	configs.LoadEnv()
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.Validator = helpers.GetValidator()

	database.ConnectDB()
	defer database.DisConnectDB()

	routes.UserRoute(e)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &echo.Map{"data": "Hello, World!"})
	})

	e.Logger.Fatal(e.Start(configs.EnvAppPort()))
}
