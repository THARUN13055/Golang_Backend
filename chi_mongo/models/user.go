package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id    primitive.ObjectID `jbson: "_id,omitempty"`
	Name  string             `bson: "name,omitempty"`
	Email string             `bson: "email,omitempty"`
}
