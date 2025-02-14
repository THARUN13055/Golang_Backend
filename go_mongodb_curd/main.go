package main

import (
	"net/http"
	_ "net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
  r.POST("/user", uc.CreateUser)
  r.DELETE("/user/:id", uc.DeleteUser)
  http.ListenAndServe(":8080",r)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb+srv://tharun:password@go.vpv6p.mongodb.net/?retryWrites=true&w=majority&appName=go")
	if err != nil {
		panic(err)
	}
	return session

}
