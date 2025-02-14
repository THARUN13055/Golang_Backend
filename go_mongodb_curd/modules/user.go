package modules

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id    bson.ObjectId `json "id" bson: "_id"`
	Name  string        `json:"name" bson: "name"`
	Email string        `json:"email" bson: "email"`
}


func getSession() *mgo.Session {
	session , err := mgo.Dial("")
	if err != nil {
		
	}
}
