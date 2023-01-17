package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-user_api_example/configs/database"
	"go-user_api_example/helpers"
	commonResponse "go-user_api_example/modules/common/response"
	"go-user_api_example/modules/user/model"
	"go-user_api_example/modules/user/repository"
	"go-user_api_example/modules/user/response"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type UserHandler struct {
	urCase repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{
		urCase: userRepository,
	}
}

func getUserCollection() *mongo.Collection {
	return database.GetCollection(database.GetDB(), "users")
}

func (uh UserHandler) CreateUser(c echo.Context) error {
	user := &model.User{}

	if err := c.Bind(&user); err != nil {
		return commonResponse.ErrorResponseJson(http.StatusUnprocessableEntity, &echo.Map{"error": err.Error()}, "error", c)
	}

	if err := c.Validate(user); err != nil {
		return commonResponse.ErrorResponseJson(http.StatusBadRequest, &echo.Map{"error": err.Error()}, "error", c)
	}

	tx, _ := uh.urCase.FindByEmail(user.Email)
	if tx != nil {
		return commonResponse.ErrorResponseJson(http.StatusConflict, &echo.Map{}, "The email has already been taken", c)
	}

	hash, _ := helpers.HashPassword(user.Password)
	user.Password = hash

	err := uh.urCase.Save(user)

	if err != nil {
		return commonResponse.ErrorResponseJson(http.StatusInternalServerError, &echo.Map{"error": err.Error()}, "error", c)
	}

	ud, _ := uh.urCase.FindByEmail(user.Email)

	return response.UserResponseJson(http.StatusCreated, ud, "success", c)
}

func (uh UserHandler) GetUser(c echo.Context) error {
	userID := c.Param("userID")

	user, err := uh.urCase.FindByID(userID)

	if err != nil {
		fmt.Println(err)
		return commonResponse.ErrorResponseJson(http.StatusNotFound, &echo.Map{"error": err.Error()}, "no content", c)
	}

	return response.UserResponseJson(http.StatusOK, user, "success", c)
}

func (uh UserHandler) EditUser(c echo.Context) error {
	user := &model.User{}
	userID := c.Param("userID")

	if err := c.Bind(&user); err != nil {
		return commonResponse.ErrorResponseJson(http.StatusUnprocessableEntity, &echo.Map{"error": err.Error()}, "error", c)
	}

	if err := c.Validate(&user); err != nil {
		return commonResponse.ErrorResponseJson(http.StatusBadRequest, &echo.Map{"error": err.Error()}, "error", c)
	}

	err := uh.urCase.Update(userID, user)

	if err != nil {
		fmt.Println(err)
		return commonResponse.ErrorResponseJson(http.StatusNotFound, &echo.Map{"error": err.Error()}, "user update failed", c)
	}

	return response.UserResponseJson(http.StatusOK, nil, "success", c)
}

func (uh UserHandler) DeleteUser(c echo.Context) error {
	userID := c.Param("userID")

	err := uh.urCase.Delete(userID)

	if err != nil {
		fmt.Println(err)
		return commonResponse.ErrorResponseJson(http.StatusNotFound, &echo.Map{"error": err.Error()}, "user delete failed", c)
	}

	return response.UserResponseJson(http.StatusOK, nil, "success", c)
}

func (uh UserHandler) GetAllUser(c echo.Context) error {
	users, err := uh.urCase.FindAll()

	if err != nil {
		fmt.Println(err)
		return commonResponse.ErrorResponseJson(http.StatusNoContent, &echo.Map{"error": err.Error()}, "no content", c)
	}

	return response.UsersResponseJson(http.StatusOK, &users, "success", c)
}
