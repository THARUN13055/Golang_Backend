package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRequest struct {
	Id    primitive.ObjectID `jbson: "_id,omitempty"`
	Name  string             `bson: "name,omitempty"`
	Email string             `bson: "email,omitempty"`
}


type UserResponse struct {
	Data interface{} `json: "data,omitempty"`
	Error string `json: "error,omitempty`
}