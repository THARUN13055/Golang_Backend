package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB_Connection() *mongo.Client {
	cxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(cxt, options.Client().ApplyURI("mongodb://username:password@localhost:27017"))
	if err != nil {
		log.Panic(err)
	}

	if err = client.Ping(cxt, nil); err != nil {
		log.Println("Db is not Pinging", err)
		return nil
	}

	fmt.Println("Database is Connected")

	return client

}

var Client *mongo.Client = DB_Connection()

var UserCollection *mongo.Collection = Client.Database("BookStoreManagement").Collection("UserCollection")
var BookCollection *mongo.Collection = Client.Database("BookStoreManagement").Collection("BookCollection")
var OrderCollection *mongo.Collection = Client.Database("BookStoreManagement").Collection("OrderCollection")
