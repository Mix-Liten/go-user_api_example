package responses

import "github.com/labstack/echo/v4"

type UserResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message,omitempty"`
	Data    *echo.Map `json:"data"`
}
