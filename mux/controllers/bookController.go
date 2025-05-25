package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tharun13055/db"
	"tharun13055/models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Now we need to get the request and seprate the id from the request
	vars := mux.Vars(r)
	bookIDstr := vars["id"]

	// now we need to find the bookid
	var bookid models.Book
	id, err := primitive.ObjectIDFromHex(bookIDstr)
	if err != nil {
		w.Write([]byte("Cant able to convert objectid form hex"))
		return
	}

	err = db.BookCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&bookid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Getting id is not founded"))
		fmt.Println(err)
		return
	}

	marshalbook, err := json.Marshal(bookid)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(marshalbook)

}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Now we need to show all the books
	cursor, err := db.BookCollection.Find(ctx, bson.D{})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("There is not data"))
		return
	}

	// now cursor contain all the data
	var listofbook models.Books
	for cursor.Next(ctx) {
		var book models.Book
		err := cursor.Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		listofbook.AddBook(book)
	}

	encData, err := json.Marshal(listofbook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while marshal the data"))
		return
	}

	w.WriteHeader(http.StatusFound)
	w.Header().Set("Context-Type", "application/json")
	w.Write(encData)

}

func AddBook(w http.ResponseWriter, r *http.Request) {
	// Create Context to session time out when db get error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var newbook models.Book

	// now we need to decode the getting book
	err := json.NewDecoder(r.Body).Decode(&newbook)
	if err != nil {
		http.Error(w, "Getting data is error", http.StatusBadRequest)
		return
	}

	// Now we need to check wether the book is already exists or not
	existBook := db.BookCollection.FindOne(ctx, bson.M{"book_name": newbook.Book_name})
	if existBook.Err() == nil {
		w.Write([]byte("Book is already exists"))
		return
	}

	// Now we need to add the book
	newbooktoadd := models.Book{
		Id:         primitive.NewObjectID(),
		Book_name:  newbook.Book_name,
		Author:     newbook.Author,
		Created_at: time.Now(),
	}

	_, err = db.BookCollection.InsertOne(ctx, newbooktoadd)
	if err != nil {
		// Handle duplicate key or other database errors
		if mongo.IsDuplicateKeyError(err) {
			http.Error(w, "data is exists", http.StatusConflict)
		} else {
			http.Error(w, "Failed to add the newbook", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User created successfully"}`))

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// create context

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// For delete the specific book we need to get the id from the request
	vars := mux.Vars(r)
	bookIdstr := vars["id"]

	// converti to objectid
	id, err := primitive.ObjectIDFromHex(bookIdstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cant able to convert the book id"))
		return
	}

	// now we need to delete the given id
	_, err = db.BookCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("There data is not founded"))
		return
	}
	w.Write([]byte("The data successfully deleted"))

}
