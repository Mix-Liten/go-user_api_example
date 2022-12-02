package repository

import (
	"go-user_api_example/modules/user/model"
)

type UserRepository interface {
	Save(*model.User) error
	Update(string, *model.User) error
	Delete(string) error
	FindByID(string) (*model.UserPublic, error)
	FindByEmail(string) (*model.UserPublic, error)
	FindAll() (model.Users, error)
}
