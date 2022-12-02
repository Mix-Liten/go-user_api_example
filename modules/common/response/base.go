package response

import "github.com/labstack/echo/v4"

type BaseResponse struct {
	Code    int       `json:"code"`
	Data    *echo.Map `json:"data"`
	Message string    `json:"message,omitempty"`
	Success bool      `json:"success"`
}

func BaseResponseJson(code int, data *echo.Map, message string, success bool, c echo.Context) error {
	br := &BaseResponse{
		Code:    code,
		Data:    data,
		Message: message,
		Success: success,
	}

	return c.JSON(code, br)
}
