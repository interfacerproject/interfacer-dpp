package database

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
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
		mongoURI := os.Getenv("MONGODB_URI")
		if mongoURI == "" {
			mongoURI = "mongodb://localhost:27017"
		}
		clientOptions := options.Client().ApplyURI(mongoURI)
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

func EnsureIndexes(client *mongo.Client) {
	collection := GetCollection(client)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "productId", Value: 1}}},
		{Keys: bson.D{{Key: "createdBy", Value: 1}}},
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "createdAt", Value: -1}}},
		{
			Keys: bson.D{{Key: "productId", Value: 1}, {Key: "batchId", Value: 1}},
			Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
				"batchId": bson.M{"$exists": true, "$ne": ""},
			}),
		},
	}

	_, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Printf("Warning: failed to create indexes: %v", err)
	} else {
		log.Println("MongoDB indexes ensured")
	}
}
