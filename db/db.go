package db

import (
	"context"
	"go-mongodb-api/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

func ConnectMongoDB() (*mongo.Client, error) {
	if clientInstance != nil {
		return clientInstance, nil
	}

	uri := config.GetConfig("MONGO_URI")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	clientInstance = client
	return client, nil
}

func GetCollection(dbName, collectionName string) (*mongo.Collection, error) {
	client, err := ConnectMongoDB()
	if err != nil {
		return nil, err
	}
	return client.Database(dbName).Collection(collectionName), nil
}
