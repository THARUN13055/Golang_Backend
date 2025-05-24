package middleware

import (
	"net/http"
	"tharun13055/data"
)

func ValidateUser(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("name")
		password := r.Header.Get("password")

		if username == "" {
			http.Error(w, "the username is empty", http.StatusBadRequest)
			return
		}
		if data.User[username] != password {
			http.Error(w, "the password is wrong", http.StatusUnauthorized)
			return
		}

		f(w, r)

	}
}

func ValidateOwner(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("name")
		password := r.Header.Get("password")
		IfOwner := r.Header.Get("userType")

		if IfOwner != "owner" {
			w.Write([]byte("you are not owner"))
			return
		}

		if username == "" {
			http.Error(w, "the username is empty", http.StatusBadRequest)
			return
		}
		if data.User[username] != password {
			http.Error(w, "the password is wrong", http.StatusUnauthorized)
			return
		}

		f(w, r)

	}
}

func TrackNumberOfReq(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data.NumberofReq++
		f(w, r)
	}
}
