package routes

import (
	"tharun13055/controllers"
	"tharun13055/middleware"

	"github.com/gorilla/mux"
)

func OrderRoutes(router *mux.Router) {
	router.HandleFunc("/orderbook/{id}", middleware.ValidateUser(controllers.OrderBooks))
	router.HandleFunc("/getorders", middleware.ValidateUser(controllers.GetAllOrders))
}
