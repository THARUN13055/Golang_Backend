package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"tharun13055/mongodb_curd/controllers"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	r := httprouter.New()
	collection := getCollection()
	uc := controllers.NewUser Controller(collection)

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":8080", r)
}

// func getSession() *mgo.Session {
// 	connection_db_url := "mongodb://tharun:password12345@3.7.247.50:27017/tharun"
// 	session, err := mgo.Dial(connection_db_url)
// 	if err != nil {
// 		fmt.Println("mongodb is not connect")
// 		panic(err)
// 	}
// 	return session
// }

func getCollection() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://tharun:password12345@3.7.247.50:27017/tharun")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("MongoDB connection error:", err)
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("MongoDB ping error:", err)
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client.Database("mongo-golang").Collection("users")
}