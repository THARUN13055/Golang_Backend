package utils

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://tharun:abcd@localhost:27017/curd?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
		return nil, err
	}

	// Check connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return client, nil
}
