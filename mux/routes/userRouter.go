package routes

import (
	"tharun13055/controllers"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/signup", controllers.Signup)
}
