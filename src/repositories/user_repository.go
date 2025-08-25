package repositories

import (
	"context"
	"go-fcontrol-api/src/configs"
	"go-fcontrol-api/src/models"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		collection: configs.MongoDatabase.Collection("users"),
	}
}

func (r *UserRepository) GetUsers(ctx context.Context, filter bson.M) ([]models.UserResponse, error) {
	slog.Debug("Searching users", "filter", filter)
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	users := make([]models.UserResponse, 0)
	if err := cursor.All(ctx, &users); err != nil {
		return make([]models.UserResponse, 0), err
	}

	return users, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	slog.Debug("Creating user", "email", user.Email)
	res, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		slog.Debug("Error to create user", "error", err)
		return nil, err
	}

	return res, nil
}
