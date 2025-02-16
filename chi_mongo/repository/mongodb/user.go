package mongodb

import (
	"tharun13055/chi_mongo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCLient struct {
	Client mongo.Collection
}

func (c MongoCLient) CreateUser(user models.User) (string, error) {
	return "", nil
}

func (c MongoCLient) GetUserId(id string) (models.User, error) {
	return models.User{}, nil
}

func (c MongoCLient) UpdateUserName(id string, name string) (int, error) {
	return 0, nil
}

func (c MongoCLient) GetAllUser() ([]models.User, error) {
	return []models.User{}, nil
}

func (c MongoCLient) DeleteUserById(id string) (int, error) {
	return 0, nil
}

func (c MongoCLient) DeleteAllUser() (int, error) {
	return 0, nil
}
