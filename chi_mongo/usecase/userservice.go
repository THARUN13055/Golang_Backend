package usecase

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"tharun13055/chi_mongo/dto"
	"tharun13055/chi_mongo/models"
	"tharun13055/chi_mongo/repository"
)

type UserService struct {
	DBClient repository.UserInterface
}

func (svc UserService) CreateUser(w http.ResponseWriter, r *http.Request)     {
	res := dto.UserResponse{}
	// extract body from the reqquest
	var UserRequest dto.UserRequest
	err := json.NewDecoder(r.Body).Decode(&UserRequest)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "invalid Request"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	user := models.User {
		Id: UserRequest.Id,
		Name: UserRequest.Name,
	    Email: UserRequest.Email,
	}

	result , err := svc.DBClient.CreateUser(user)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "Error while inserting the data"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("user successfully created", slog.String("_id",result))
	res.Data = result
	json.NewEncoder(w).Encode(res)
}
func (svc UserService) GetUserId(w http.ResponseWriter, r *http.Request)      {}
func (svc UserService) UpdateUserName(w http.ResponseWriter, r *http.Request) {}
func (svc UserService) GetAllUser(w http.ResponseWriter, r *http.Request)     {}
func (svc UserService) DeleteUserById(w http.ResponseWriter, r *http.Request) {}
func (svc UserService) DeleteAllUser(w http.ResponseWriter, r *http.Request)  {}
