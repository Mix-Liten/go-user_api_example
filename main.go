package main

import (
	"github.com/labstack/echo/v4"
	"go-user_api_example/configs"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &echo.Map{"data": "Hello, World!"})
	})
	e.Logger.Fatal(e.Start(configs.ENVAPPPort()))
}
