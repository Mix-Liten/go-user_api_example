package responses

import "github.com/labstack/echo/v4"

type UserResponse struct {
	BaseResponse
	Data *echo.Map `json:"data"`
}

//type createUser struct {
//	FirstName string   `json:"firstname"`
//	LastName  string   `json:"lastname"`
//	Email     string   `json:"email"`
//	Birth     string   `json:"birth"`
//	Languages []string `json:"languages"`
//}
