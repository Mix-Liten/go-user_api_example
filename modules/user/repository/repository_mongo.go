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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)

	return err
}

func (r *userRepositoryMongo) Update(userID string, user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.UpdatedAt = time.Now()
	objId, _ := primitive.ObjectIDFromHex(userID)

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": objId}, user)

	return err
}

func (r *userRepositoryMongo) Delete(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userID)

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": objId})

	return err
}

func (r *userRepositoryMongo) FindByID(userID string) (*model.UserPublic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user model.UserPublic

	objId, _ := primitive.ObjectIDFromHex(userID)

	err := r.collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryMongo) FindByEmail(userEmail string) (*model.UserPublic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user model.UserPublic

	err := r.collection.FindOne(ctx, bson.M{"email": userEmail}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryMongo) FindAll() (model.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users model.Users

	results, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser model.UserPublic
		if err = results.Decode(&singleUser); err != nil {
			return nil, err
		}

		users = append(users, singleUser)
	}

	return users, nil
}
