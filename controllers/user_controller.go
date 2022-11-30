package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-user_api_example/configs"
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

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func (uc UserController) CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	user := &models.UserModel{}
	defer cancel()

	if err := c.Bind(&user); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &echo.Map{"data": err.Error()},
			})
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &echo.Map{"data": err.Error()},
			})
	}

	tx, _ := findUser(user)
	if tx.Email != "" {
		return c.JSON(
			http.StatusBadRequest,
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "The email has already been taken",
				Data:    &echo.Map{},
			})
	}

	hash, _ := helpers.HashPassword(user.Password)
	user.Password = hash

	result, err := userCollection.InsertOne(ctx, user)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &echo.Map{"data": err.Error()},
			})
	}

	return c.JSON(
		http.StatusCreated,
		responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &echo.Map{"data": result},
		})
}

func findUser(user *models.UserModel) (models.UserModel, error) {
	filter := bson.M{"email": user.Email}
	result := models.UserModel{}
	err := userCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.UserModel{}, nil
		}

		return models.UserModel{}, fmt.Errorf("finding user: %w", err)
	}

	return result, nil
}
