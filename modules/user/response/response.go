package response

import (
	"github.com/labstack/echo/v4"
	"go-user_api_example/helpers"
	"go-user_api_example/modules/common/response"
	"go-user_api_example/modules/user/model"
)

func UserResponseJson(code int, data *model.UserPublic, message string, c echo.Context) error {
	if data == nil {
		return response.BaseResponseJson(code, nil, message, true, c)
	}

	up := helpers.StructToMap(&model.UserPublic{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Birth:     data.Birth,
		Languages: data.Languages,
	})

	return response.BaseResponseJson(code, up, message, true, c)
}

func UsersResponseJson(code int, data *model.Users, message string, c echo.Context) error {
	us := &echo.Map{"users": *data}

	return response.BaseResponseJson(code, us, message, true, c)
}
