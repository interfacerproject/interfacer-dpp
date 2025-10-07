package database

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

const (
	DBName         = "dpp_db"
	CollectionName = "passports"
)

func ConnectDB() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(clientOptions)
		if err != nil {
			clientInstanceError = err
			return
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			clientInstanceError = err
			return
		}

		log.Println("Successfully connected to MongoDB!")
		clientInstance = client
	})

	return clientInstance, clientInstanceError
}

func GetCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database(DBName).Collection(CollectionName)
	return collection
}