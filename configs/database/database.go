package database

import (
	"context"
	"fmt"
	"go-user_api_example/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var db *mongo.Client
var cc *context.Context

func ConnectDB() *mongo.Client {
	credential := options.Credential{
		Username: configs.EnvDBUsername(),
		Password: configs.EnvDBPassword(),
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.EnvMongoURI()).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(fmt.Sprintf("db connect error: %s", err.Error()))
	}

	// ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("db ping error: %s", err.Error()))
	}

	db = client
	cc = &ctx

	fmt.Println("Connected to MongoDB")

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(configs.EnvProjectName()).Collection(collectionName)
	return collection
}

func GetDB() *mongo.Client {
	return db
}

func DisConnectDB() error {
	err := db.Disconnect(*cc)
	return err
}
