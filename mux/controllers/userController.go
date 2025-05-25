package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tharun13055/db"
	"tharun13055/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Create Context to session time out when db get error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// You are passing the json request body to struct
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Getting User data error", http.StatusBadRequest)
		return
	}

	// Here we need to check user.Name is empty or not to avoid config
	if user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The username is empty"))
		fmt.Print(err)
		return
	}

	// Now after convert into struct we need to check with db
	var user_id_db models.User
	err = db.UserCollection.FindOne(ctx, bson.M{"name": user.Name}).Decode(&user_id_db)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Body reading getting error"))
		return
	}

	// Now if we founded we need to check the password is correct or not.
	if user_id_db.Password != user.Password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}

func Signup(w http.ResponseWriter, r *http.Request) {
	// Create Context to session time out when db get error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Now we need to create the struct
	var newuser models.User
	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		http.Error(w, "Getting json data is error", http.StatusBadRequest)
		return
	}

	// New we need to Find that the user is already in db or not
	var existingUser models.User
	err = db.UserCollection.FindOne(ctx, bson.M{"name": newuser.Name}).Decode(&existingUser)
	if err == nil {
		// User found
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	// Adding the new user in the db
	newuserInsertDb := models.User{
		Id:         primitive.NewObjectID(),
		Name:       newuser.Name,
		Password:   newuser.Password,
		UserType:   newuser.UserType,
		Created_at: time.Now(),
	}

	// Insert into the db
	_, err = db.UserCollection.InsertOne(ctx, newuserInsertDb)
	if err != nil {
		// Handle duplicate key or other database errors
		if mongo.IsDuplicateKeyError(err) {
			http.Error(w, "User already exists", http.StatusConflict)
		} else {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			fmt.Print(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User created successfully"}`))

}
