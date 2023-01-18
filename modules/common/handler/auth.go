package handler

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go-user_api_example/configs/database"
	"go-user_api_example/helpers"
	"go-user_api_example/modules/user/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"strconv"
	"time"
)

type TokenClaims struct {
	UserID primitive.ObjectID
	jwt.StandardClaims
}

type CreateTokenRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func CreateToken(c echo.Context) error {
	req := &CreateTokenRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	ur := repository.NewUserRepositoryMongo(database.GetDB(), "users")
	user, err := ur.FindByEmail(req.Email)
	if err != nil {
		return echo.ErrUnauthorized
	}
	if !helpers.CheckPassword(req.Password, user.Password) {
		return echo.ErrUnauthorized
	}

	ttl, err := strconv.Atoi(os.Getenv("JWT_TTL"))
	if err != nil {
		return err
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
		},
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
