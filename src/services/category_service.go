package services

import (
	"context"
	"go-fcontrol-api/src/models"
	"go-fcontrol-api/src/repositories"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CategoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		repo: *repositories.NewCategoryRepository(),
	}
}

func (s *CategoryService) GetCategory(ctx context.Context, queryParam models.Category) ([]models.Category, error) {
	filter := bson.M{}

	if queryParam.Name != "" || queryParam.Type != "" {
		filter = bson.M{
			"$and": bson.A{
				bson.M{"name": bson.M{"$regex": queryParam.Name, "$options": "i"}},
				bson.M{"type": bson.M{"$regex": queryParam.Type, "$options": "i"}},
			},
		}
	}

	return s.repo.GetCategory(ctx, filter)
}

func (s *CategoryService) GetById(ctx context.Context, id string) (models.Category, error) {
	category, err := s.repo.FindById(ctx, id)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (s *CategoryService) CreateCategory(ctx context.Context, category models.Category) (*models.Category, error) {
	res, err := s.repo.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	slog.Info("Category created", "id", res.InsertedID)

	category.ID = res.InsertedID.(bson.ObjectID)

	return &category, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, id string, input models.Category) (models.Category, error) {
	_, err := s.repo.FindById(ctx, id)
	if err != nil {
		slog.Debug("Category not found", "error", err)
		return models.Category{}, err
	}

	_, err = s.repo.UpdateCategory(ctx, id, input)
	if err != nil {
		return models.Category{}, err
	}

	updatedCategory, err := s.repo.FindById(ctx, id)
	if err != nil {
		return models.Category{}, err
	}

	return updatedCategory, nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id string) (models.Category, error) {
	slog.Debug("Removing category", "id", id)

	category, err := s.repo.FindById(ctx, id)
	if err != nil {
		slog.Debug("Category not found", "error", err)
		return models.Category{}, err
	}

	_, err = s.repo.DeleteCategory(ctx, category.ID.Hex())
	if err != nil {
		return models.Category{}, err
	}

	return models.Category{}, nil
}
