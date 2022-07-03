package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDBConnection *mongo.Database

func connect() {
	credentialOpts := options.Credential{
		Username: "root",
		Password: "example",
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017").SetAuth(credentialOpts)
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
