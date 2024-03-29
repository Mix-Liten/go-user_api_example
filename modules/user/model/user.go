package model

import (
	"go-user_api_example/modules/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	Password  string             `json:"password" bson:"password" validate:"required,min=8"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	Birth     string             `json:"birth" bson:"birth"`
	Languages []string           `json:"languages" bson:"languages"`
	model.Base
}

type UserPublic struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty""`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Email     string             `json:"email"`
	Birth     string             `json:"birth"`
	Languages []string           `json:"languages"`
}

type Users []UserPublic
