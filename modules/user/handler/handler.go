package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-user_api_example/configs/database"
	"go-user_api_example/helpers"
	commonResponse "go-user_api_example/modules/common/response"
	"go-user_api_example/modules/user/model"
	"go-user_api_example/modules/user/repository"
	"go-user_api_example/modules/user/response"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type UserHandler struct {
	ur repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{
		ur: userRepository,
	}
}

func getUserCollection() *mongo.Collection {
	return database.GetCollection(database.GetDB(), "users")
}

func (uh UserHandler) CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	user := &model.User{}
	defer cancel()

	if err := c.Bind(&user); err != nil {
		return commonResponse.ErrorResponseJson(http.StatusUnprocessableEntity, &echo.Map{"error": err.Error()}, "error", c)
	}

	if err := c.Validate(&user); err != nil {
		return commonResponse.ErrorResponseJson(http.StatusBadRequest, &echo.Map{"error": err.Error()}, "error", c)
	}

	tx, _ := uh.ur.FindByEmail(user.Email)
	if tx.Email != "" {
		return commonResponse.ErrorResponseJson(http.StatusConflict, &echo.Map{}, "The email has already been taken", c)
	}

	hash, _ := helpers.HashPassword(user.Password)
	user.Password = hash

	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	_, err := getUserCollection().InsertOne(ctx, user)

	if err != nil {
		return commonResponse.ErrorResponseJson(http.StatusInternalServerError, &echo.Map{"error": err.Error()}, "error", c)
	}

	ud, _ := uh.ur.FindByEmail(user.Email)

	return response.UserResponseJson(http.StatusCreated, ud, "success", c)
}
