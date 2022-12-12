package repository

import (
	"context"
	"go-user_api_example/configs/database"
	"go-user_api_example/modules/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user model.UserPublic
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userID)

	err := r.collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryMongo) FindByEmail(userEmail string) (*model.UserPublic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user model.UserPublic
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"email": userEmail}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryMongo) FindAll() (model.Users, error) {
	return nil, nil
}
