package main

import (
	"net/http"
	"tharun13055/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.UserRoutes(router)
	routes.BookRoutes(router)
	routes.OrderRoutes(router)

	http.ListenAndServe(":3030", router)
}
