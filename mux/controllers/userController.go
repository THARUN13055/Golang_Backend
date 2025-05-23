package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tharun13055/data"
	"tharun13055/datatypes"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var user datatypes.UserCredential

	// This error is used to verify the user credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while decoding the user credentials", http.StatusBadRequest)
		return
	}

	// Check if the user is existing or not
	_, exists := data.User[user.Name]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	// if User is correct then we need to check the password

	if data.User[user.Name] != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid password"))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
		return
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var newUser datatypes.UserCredential

	// check wether the use is already existing or not
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while decoding the user credentials", http.StatusBadRequest)
		return
	}

	_, exists := data.User[newUser.Name]
	if exists {
		w.Write([]byte("User already exists"))
		return
	}
	// if user is not existing then we need to add the user to the map

	data.User[newUser.Name] = newUser.Password
	w.Write([]byte("New user created successfully"))
	w.WriteHeader(http.StatusCreated)
}
