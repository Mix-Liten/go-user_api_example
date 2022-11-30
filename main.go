package main

import (
	"github.com/labstack/echo/v4"
	"go-user_api_example/configs"
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
	e.Validator = helpers.GetValidator()

	configs.ConnectDB()
	defer configs.DisConnectDB()

	routes.UserRoute(e)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &echo.Map{"data": "Hello, World!"})
	})
	e.Logger.Fatal(e.Start(configs.EnvAppPort()))
}
