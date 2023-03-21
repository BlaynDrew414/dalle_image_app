package db

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

// ConnectToDB creates a connection to the MongoDB database
func ConnectToDB(collectionName string) (*mongo.Collection, error) {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        return nil, fmt.Errorf("MONGODB_URI environment variable not set")
    }

    clientOptions := options.Client().ApplyURI(mongoURI)

    // connect to db
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return nil, err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // check connection
    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err
    }

    // select database and collection
    db := client.Database("dalle_image_app")
    collection := db.Collection(collectionName)

    return collection, nil
}



