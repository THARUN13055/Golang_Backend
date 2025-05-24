package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tharun13055/data"
	"tharun13055/datatypes"

	"github.com/gorilla/mux"
)

func OrderBooks(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("name")

	// Here we are getting the url of the book
	vars := mux.Vars(r)
	bookIDstr, errbool := vars["id"]
	if !errbool {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Here we are changing the ID as string to int
	bookId, err := strconv.Atoi(bookIDstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cant able to change string into int"))
		return
	}

	// we need to check wether there is book or not
	_, exists := data.Book[bookId]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("All the book has been solded"))
		return
	}

	// Now we need to add the order
	data.Order[username] = append(data.Order[username], bookId)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Book added successfully"))
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("name")

	var listofbook datatypes.Books

	orders := data.Order[username]

	for _, bookId := range orders {
		book := datatypes.Book{
			Id:   bookId,
			Name: data.Book[bookId],
		}
		listofbook.AddBook(book)
	}

	// Which we need to view so we are making as marshal
	marshaledData, err := json.Marshal(listofbook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ordered list or marshal getting error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(marshaledData)
}
