package routes

import (
	"tharun13055/controllers"

	"github.com/gorilla/mux"
)

func OrderRoutes(router *mux.Router) {
	router.HandleFunc("/orderbooks", controllers.OrderBooks)
	router.HandleFunc("/getorders", controllers.GetAllOrders)
}
