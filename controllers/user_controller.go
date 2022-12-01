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

//var userCollection *mongo.Collection = database.GetCollection(database.GetDB(), "users")

func (uc UserController) CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	user := &models.UserModel{}
	defer cancel()

	er := &responses.ErrorResponse{
		Data: &echo.Map{},
	}

	if err := c.Bind(&user); err != nil {
		er.Status = http.StatusBadRequest
		er.Message = "error"
		er.Data = &echo.Map{"error": err.Error()}
		return c.JSON(http.StatusBadRequest, er)
	}

	if err := c.Validate(user); err != nil {
		er.Status = http.StatusBadRequest
		er.Message = "error"
		er.Data = &echo.Map{"error": err.Error()}
		return c.JSON(http.StatusBadRequest, er)
	}

	tx, _ := findUser(user)
	if tx.Email != "" {
		er.Status = http.StatusBadRequest
		er.Message = "The email has already been taken"
		return c.JSON(http.StatusBadRequest, er)
	}

	hash, _ := helpers.HashPassword(user.Password)
	user.Password = hash

	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	_, err := getUserCollection().InsertOne(ctx, user)

	if err != nil {
		er.Status = http.StatusInternalServerError
		er.Message = "error"
		er.Data = &echo.Map{"error": err.Error()}
		return c.JSON(http.StatusInternalServerError, er)
	}

	ud, _ := findUser(user)

	ur := &responses.UserResponse{}
	ur.Status = http.StatusCreated
	ur.Message = "success"
	ur.Data = &echo.Map{"user": &echo.Map{
		"firstName": ud.FirstName,
		"lastName":  ud.LastName,
		"email":     ud.Email,
		"birth":     ud.Birth,
		"languages": ud.Languages,
	}}

	return c.JSON(http.StatusCreated, ur)
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
