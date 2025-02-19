package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Retrieve environment variables
	dbName := os.Getenv("MONGO_DB_NAME")
	dbPassword := os.Getenv("MONGO_DB_PASSWORD")
	dbURI := os.Getenv("MONGO_URI")
	dbPort := os.Getenv("MONGO_DB_PORT")

	// Construct the MongoDB connection URI
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", dbName, dbPassword, dbURI, dbPort, dbName)

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// sending the ping that it is working or not

	if err := client.Database(os.Getenv("MONGO_DB_NAME")).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Database is successfully connected")

	return client

}
