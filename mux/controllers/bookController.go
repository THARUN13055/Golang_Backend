package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tharun13055/data"
	"tharun13055/datatypes"

	"github.com/gorilla/mux"
)

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	// Now we need to get the request and seprate the id from the request
	vars := mux.Vars(r)
	bookIDstr := vars["id"]

	// Now we need to convert id into string
	bookID, err := strconv.Atoi(bookIDstr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusInternalServerError)
		w.Write([]byte("Invalid book ID"))
		return
	}
	// Now using this id we need to get all the details of the book
	RequiredBook, bookexits := data.Book[bookID]
	if !bookexits {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found"))
		return
	}

	// Now if book is founded then we need to marshal the data into json
	// until now it is in json format

	marshalBook, err := json.Marshal(RequiredBook)
	if err != nil {
		http.Error(w, "Error while marshalling the book data", http.StatusInternalServerError)
		w.Write([]byte("Error while marshalling the book data"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalBook)

}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Now we need to show all the books
	listofBooks := data.Book

	// this we are adding the list of book to show it
	var showlistofbook datatypes.Books
	// loop all the books and add to the list

	for bid, bookname := range listofBooks {
		book := datatypes.Book{
			Id:   bid,
			Name: bookname,
		}
		showlistofbook.AddBook(book)
	}

	// Now we need to marshal the data into json

	marshalBook, err := json.Marshal(showlistofbook)
	if err != nil {
		http.Error(w, "Error while marshalling the book data", http.StatusInternalServerError)
		w.Write([]byte("Error while marshalling the book data"))
		return
	}

	// if all success then we need to send the data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalBook)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// For delete the specific book we need to get the id from the request
	vars := mux.Vars(r)
	bookIdstr := vars["id"]

	// Now we need to convert the id into int

	bookID, err := strconv.Atoi(bookIdstr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusInternalServerError)
		w.Write([]byte("Invalid Book id"))
		return
	}

	// Now we need to check wether the book is exists or not

	_, exists := data.Book[bookID]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found"))
		return
	}

	// If there is Book id we need to delete the book

	delete(data.Book, bookID)
	w.Write([]byte("Book deleted successfully"))

}

func AddBook(w http.ResponseWriter, r *http.Request) {

}
