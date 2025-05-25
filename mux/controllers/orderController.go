package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"tharun13055/db"
	"tharun13055/models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func OrderBooks(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// to verify wether the user is login or not
	username := r.Header.Get("username")

	// Here we are getting the url of the book
	vars := mux.Vars(r)

	bookIdstr, errbool := vars["book_id"]
	if !errbool {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("id not founded"))
		return
	}

	// We need to change the id into the primitive data
	bookid, err := primitive.ObjectIDFromHex(bookIdstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("here objectid can't convert into hex"))
		return
	}

	// Now we need to check book id is exists or not
	bookexist := db.BookCollection.FindOne(ctx, bson.M{"book_id": bookid})
	if bookexist.Err() != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Here we cant able to find the book"))
		return
	}

	order_Book := models.Order{
		Id:         primitive.NewObjectID(),
		User_id:    username,
		Book_id:    bookIdstr,
		Created_at: time.Now(),
	}

	_, err = db.OrderCollection.InsertOne(ctx, order_Book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Set status for internal server error
		w.Write([]byte("Failed to add the order"))
		return
	}

	w.WriteHeader(http.StatusOK) // Set the status code for a successful operation
	w.Write([]byte("Added the data successfully"))
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	username := r.Header.Get("name")

	// Here we need to use cursor to list the orders
	cursor, err := db.OrderCollection.Find(ctx, bson.M{"user_id": username})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Here no Order is fouunded"))
		return
	}
	defer cursor.Close(ctx)

	var listofOrder models.Orders

	for cursor.Next(ctx) {
		var order models.Order
		err := cursor.Decode(&order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		listofOrder.AddOrder(order)
	}

	encData, err := json.Marshal(listofOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(encData)
}
