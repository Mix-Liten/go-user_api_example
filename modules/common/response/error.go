package response

import "github.com/labstack/echo/v4"

func ErrorResponseJson(code int, data *echo.Map, message string, c echo.Context) error {
	return BaseResponseJson(code, data, message, false, c)
}
