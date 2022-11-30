package responses

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	BaseResponse
	Data *echo.Map `json:"data"`
}
