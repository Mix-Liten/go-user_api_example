package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var DB *mongo.Client = ConnectDB()
var cc *context.Context

func ConnectDB() *mongo.Client {
	credential := options.Credential{
		Username: EnvDBUsername(),
		Password: EnvDBPassword(),
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()).SetAuth(credential))
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

	cc = &ctx

	fmt.Println("Connected to MongoDB")

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(EnvProjectName()).Collection(collectionName)
	return collection
}

func DisConnectDB() error {
	err := DB.Disconnect(*cc)
	return err
}
