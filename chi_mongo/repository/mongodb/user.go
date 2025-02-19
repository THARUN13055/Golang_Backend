package mongodb

import (
	"context"
	"fmt"
	"log/slog"
	"tharun13055/chi_mongo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type MongoCLient struct {
	Client mongo.Collection
}

// Create user

func (c MongoCLient) CreateUser(user models.User) (string, error) {
	result, err := c.Client.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// Get the user by id
func (c MongoCLient) GetUserId(id string) (models.User, error) {
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid id")
	}
	var user models.User
	filter := bson.D{{Name: "_id", Value: docId}}

	err = c.Client.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

// Update the user by id

func (c MongoCLient) UpdateUserName(id string, name string) (int, error) {

	docId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, fmt.Errorf("invalid id")
	}

	filter := bson.D{{Name: "_id", Value: docId}}

	updatingvalue := bson.D{{Name: "$set", Value: bson.D{{Name: "name", Value: name}}}}

	result, err := c.Client.UpdateOne(context.Background(), filter, updatingvalue)

	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil

}

func (c MongoCLient) GetAllUser() ([]models.User, error) {
	filter := bson.D{}

	getalluser, err := c.Client.Find(context.Background(), filter)

	if err != nil {
		return []models.User{}, err
	}
	defer getalluser.Close(context.Background())

	var users []models.User

	for getalluser.Next(context.Background()) {
		var user models.User
		err := getalluser.Decode(&user)
		if err != nil {
			slog.Error("Error while decoding the user", slog.String("error", err.Error()))
			continue
		}
		users = append(users, user)
	}

	return users, nil
}

func (c MongoCLient) DeleteUserById(id string) (int, error) {
	docId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, fmt.Errorf("invalid id")
	}

	filter := bson.D{{Name: "_id", Value: docId}}

	result, err := c.Client.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}

func (c MongoCLient) DeleteAllUser() (int, error) {

	filter := bson.D{}

	result, err := c.Client.DeleteMany(context.Background(), filter)

	if err != nil {
		return 0, err
	}
	return int(result.DeletedCount), nil
}
