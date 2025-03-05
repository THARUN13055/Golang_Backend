package main

import (
	"fmt"
	"log"
	"net/http"
	"tharun13055/utils"
)

func main() {
	// connect the db

	client, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("mongodb is not connect!", err)
	}

	defer client.Disconnect(nil)

	// Set up routes
	router := routes.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("server is started")
}
