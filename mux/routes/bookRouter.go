package routes

import (
	"tharun13055/controllers"

	"github.com/gorilla/mux"
)

func BookRoutes(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks)
	router.HandleFunc("/getbook/{id}", controllers.GetBookByID)
	router.HandleFunc("/addbook", controllers.AddBook)
	router.HandleFunc("/deletebook/{id}", controllers.DeleteBook)
}
