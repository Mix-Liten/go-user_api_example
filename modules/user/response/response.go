package response

import (
	"github.com/labstack/echo/v4"
	"go-user_api_example/helpers"
	"go-user_api_example/modules/common/response"
	"go-user_api_example/modules/user/model"
)

func UserResponseJson(code int, data model.User, message string, c echo.Context) error {
	up := &model.UserPublic{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Birth:     data.Birth,
		Languages: data.Languages,
	}

	return response.BaseResponseJson(code, helpers.StructToMap(up), message, true, c)
}
