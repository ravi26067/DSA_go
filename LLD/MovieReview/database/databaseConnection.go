package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// StartDB initialize the db connection
func StartDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	MongoDb := os.Getenv("MONGOURI")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB Successfully!")
	return client
}

// Client initiates the db client
var Client *mongo.Client = StartDB()

// OpenCollection used to return the collection for the given collection name
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("shive-api").Collection(collectionName)
	return collection
}
