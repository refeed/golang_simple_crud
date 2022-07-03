package db

import (
	"context"
	"fmt"
	"golangSimpleCrud/lib"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Might be a good idea to just provide MONGODB_STRING as a env variable
	// instead of these three variables
	MONGODB_HOST     = lib.GetEnvOrDefault("MONGODB_HOST", "localhost:27017")
	MONGODB_USERNAME = lib.GetEnvOrDefault("MONGODB_USERNAME", "root")
	MONGODB_PASSWORD = lib.GetEnvOrDefault("MONGODB_PASSWORD", "example")
)

var mongoDBConnection *mongo.Database

func connect() {
	credentialOpts := options.Credential{
		Username: MONGODB_USERNAME,
		Password: MONGODB_PASSWORD,
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI(fmt.Sprintf("mongodb://%v", MONGODB_HOST)).SetAuth(credentialOpts)
	client, _ := mongo.NewClient(clientOptions)
	err := client.Connect(context.TODO())
	if err != nil {
		panic(err)
	}

	mongoDBConnection = client.Database("golang_simple_crud")
}

func GetDB() *mongo.Database {
	if mongoDBConnection == nil {
		connect()
	}
	return mongoDBConnection
}
