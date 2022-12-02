package repository

import (
	"go-user_api_example/configs/database"
	"go-user_api_example/modules/user/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserRepositoryMongo(db *mongo.Client, collectionName string) *userRepositoryMongo {
	collection := database.GetCollection(db, collectionName)
	return &userRepositoryMongo{
		collection: collection,
	}
}

func (r *userRepositoryMongo) Save(user *model.User) error {
	return nil
}

func (r *userRepositoryMongo) Update(userID string, user *model.User) error {
	return nil
}

func (r *userRepositoryMongo) Delete(userID string) error {
	return nil
}

func (r *userRepositoryMongo) FindByID(userID string) (*model.UserPublic, error) {
	return nil, nil
}

func (r *userRepositoryMongo) FindByEmail(userEmail string) (*model.UserPublic, error) {
	return nil, nil
}

func (r *userRepositoryMongo) FindAll() (model.Users, error) {
	return nil, nil
}
