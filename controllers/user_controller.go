package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-user_api_example/configs/database"
	"go-user_api_example/helpers"
	"go-user_api_example/models"
	"go-user_api_example/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type UserController struct {
	context *echo.Context
}

func NewUserController() *UserController {
	return &UserController{}
}

func getUserCollection() *mongo.Collection {
	return database.GetCollection(database.GetDB(), "users")
}

func (uc UserController) CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	user := &models.UserModel{}
	defer cancel()

	if err := c.Bind(&user); err != nil {
		return responses.ErrorResponseJson(http.StatusUnprocessableEntity, &echo.Map{"error": err.Error()}, "error", c)
	}

	if err := c.Validate(&user); err != nil {
		return responses.ErrorResponseJson(http.StatusBadRequest, &echo.Map{"error": err.Error()}, "error", c)
	}

	tx, _ := findUser(user)
	if tx.Email != "" {
		return responses.ErrorResponseJson(http.StatusConflict, &echo.Map{}, "The email has already been taken", c)
	}

	hash, _ := helpers.HashPassword(user.Password)
	user.Password = hash

	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	_, err := getUserCollection().InsertOne(ctx, user)

	if err != nil {
		return responses.ErrorResponseJson(http.StatusInternalServerError, &echo.Map{"error": err.Error()}, "error", c)
	}

	ud, _ := findUser(user)

	return responses.UserResponseJson(http.StatusCreated, ud, "success", c)
}

func findUser(user *models.UserModel) (models.UserModel, error) {
	filter := bson.M{"email": user.Email}
	result := models.UserModel{}
	err := getUserCollection().FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.UserModel{}, nil
		}

		return models.UserModel{}, fmt.Errorf("finding user: %w", err)
	}

	return result, nil
}
