package services

import (
	"context"
	"go-fcontrol-api/src/models"
	"go-fcontrol-api/src/repositories"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: *repositories.NewUserRepository(),
	}
}

func (s *UserService) GetUser(ctx context.Context, queryParam models.UserResponse) ([]models.UserResponse, error) {
	filter := bson.M{}

	if queryParam.Name != "" {
		filter = bson.M{
			"$and": bson.A{
				bson.M{"name": bson.M{"$regex": queryParam.Name, "$options": "i"}},
			},
		}
	}

	return s.repo.GetUsers(ctx, filter)
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) (*models.UserResponse, error) {
	res, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	slog.Info("User created", "id", res.InsertedID, "email", user.Email)

	userResponse := models.UserResponse{
		ID:    res.InsertedID.(bson.ObjectID),
		Name:  user.Name,
		Email: user.Email,
	}

	return &userResponse, nil
}
