package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	Password  string             `json:"password,omitempty" bson:"password" validate:"required,min=8"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	Birth     string             `json:"birth" bson:"birth"`
	Languages []string           `json:"languages" bson:"languages"`
	BaseModel
}
