package usecase

import (
	"net/http"
	"tharun13055/chi_mongo/repository"
)

type UserService struct {
	DBClient repository.UserInterface
}

func (svc UserService) CreateUser(w http.ResponseWriter, r *http.Request)     {}
func (svc UserService) GetUserId(w http.ResponseWriter, r *http.Request)      {}
func (svc UserService) UpdateUserName(w http.ResponseWriter, r *http.Request) {}
func (svc UserService) GetAllUser(w http.ResponseWriter, r *http.Request)     {}
func (svc UserService) DeleteUserById(w http.ResponseWriter, r *http.Request) {}
func (svc UserService) DeleteAllUser(w http.ResponseWriter, r *http.Request)  {}
