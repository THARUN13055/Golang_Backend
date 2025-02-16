package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tharun13055/mongodb_curd/models"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserController struct
type UserController struct {
	Collection *mongo.Collection
}

// NewUser Controller creates a new UserController
func NewUser Controller(collection *mongo.Collection) *User Controller {
	return &User Controller{collection}
}

// GetUser  retrieves a user by ID
func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	var u models.User

	err := uc.Collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uj)
}

// CreateUser  creates a new user
func (uc UserController) CreateUser (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.Id = bson.NewObjectId()

	_, err := uc.Collection.InsertOne(context.TODO(), u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(uj)
}

// DeleteUser  deletes a user by ID
func (uc UserController) DeleteUser (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	_, err := uc.Collection.DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted the user with ID: %s\n", oid.Hex())
}