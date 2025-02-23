package usecase

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"tharun13055/chi_mongo/dto"
	"tharun13055/chi_mongo/models"
	"tharun13055/chi_mongo/repository"

	"github.com/go-chi/chi/v5"
)

type UserService struct {
	DBClient repository.UserInterface
}

func (svc UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	user := models.User{
		Id:    UserRequest.Id,
		Name:  UserRequest.Name,
		Email: UserRequest.Email,
	}

	result, err := svc.DBClient.CreateUser(user)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "Error while inserting the data"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("user successfully created", slog.String("_id", result))
	res.Data = result
	json.NewEncoder(w).Encode(res)
}
func (svc UserService) GetUserId(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	id := chi.URLParam(r, "id")
	if id == "" {
		slog.Error("Id field is empty")
		res.Error = "no id"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	user, err := svc.DBClient.GetUserId(id)
	if err != nil {
		slog.Error("error while fetching the user")
		res.Error = "error while fetching"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("User is fetched succuessfully")
	res.Data = user
	json.NewEncoder(w).Encode(res)

}
func (svc UserService) UpdateUserName(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	userid := chi.URLParam(r, "id")
	if userid == "" {
		slog.Error("The user id is empty")
		res.Error = "the user id is empty check it"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	// extract body from the request
	var userReq dto.UserRequest

	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "the user id is empty check it"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	//

	result, err := svc.DBClient.UpdateUserName(userid, userReq.Name)
	if err != nil {
		slog.Error("the user is not imported")
		res.Error = " the user is not imported"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("Updated successfully")
	res.Data = result
	json.NewEncoder(w).Encode(res)

}
func (svc UserService) GetAllUser(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	users, err := svc.DBClient.GetAllUser()
	if err != nil {
		slog.Error("error while fetching the user")
		res.Error = "error while fetching"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("User is fetched succuessfully")
	res.Data = users
	json.NewEncoder(w).Encode(res)
}
func (svc UserService) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	id := chi.URLParam(r, "id")
	if id == "" {
		slog.Error("Id field is empty")
		res.Error = "no id"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	user, err := svc.DBClient.DeleteUserById(id)
	if err != nil {
		slog.Error("error while fetching the user")
		res.Error = "error while fetching"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("User is Deleted succuessfully")
	res.Data = user
	json.NewEncoder(w).Encode(res)
}
func (svc UserService) DeleteAllUser(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	users, err := svc.DBClient.DeleteAllUser()
	if err != nil {
		slog.Error("error to delete all user")
		res.Error = "error while fetching"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("User Delete succuessfully")
	res.Data = users
	json.NewEncoder(w).Encode(res)
}
