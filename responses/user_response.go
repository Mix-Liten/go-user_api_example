package responses

import (
	"github.com/labstack/echo/v4"
	"go-user_api_example/helpers"
	"go-user_api_example/models"
)

type UserPublic struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Email     string   `json:"email"`
	Birth     string   `json:"birth"`
	Languages []string `json:"languages"`
}

func UserResponseJson(code int, data models.UserModel, message string, c echo.Context) error {
	up := &UserPublic{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Birth:     data.Birth,
		Languages: data.Languages,
	}

	return BaseResponseJson(code, helpers.StructToMap(up), message, true, c)
}
