package main

import (
	"log"
	"log/slog"
	"net/http"
	"tharun13055/chi_mongo/usecase"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("There is no env file")
	}

	slog.Info("env is loading successfully")
}

func main() {
	//userservice importing instance

	userService := usecase.UserService{}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("the server is healthy........"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/user", userService.CreateUser)
		r.Get("/user/{id}", userService.GetUserId)
		r.Put("/user/{id}", userService.UpdateUserName)
		r.Delete("/user/{id}", userService.DeleteUserById)
		r.Delete("/user", userService.DeleteAllUser)
	})

	http.ListenAndServe(":8080", r)

}
