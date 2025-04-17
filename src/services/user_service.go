package services

import (
	"go-fcontrol-api/src/models"
	"go-fcontrol-api/src/repositories"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetAllUsers() ([]models.UserResponse, error) {
	return repositories.GetUsers()
}

func CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	return repositories.CreateUser(user)
}
