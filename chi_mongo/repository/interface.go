package repository

import "tharun13055/chi_mongo/models"

type UserInterface interface {
	CreateUser(models.User) (string, error)
	GetUserId(string) (models.User, error)
	UpdateUserName(string, string) (int, error)
	GetAllUser() ([]models.User, error)
	DeleteUserById(string) (int, error)
	DeleteAllUser() (int, error)
}
