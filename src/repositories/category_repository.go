package repositories

import (
	"context"
	"go-fcontrol-api/src/configs"
	"go-fcontrol-api/src/models"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		collection: configs.MongoDatabase.Collection("categories"),
	}
}

func (r *CategoryRepository) GetCategory(ctx context.Context, filter bson.M) ([]models.Category, error) {
	slog.Debug("Searching categories", "filter", filter)
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	categories := make([]models.Category, 0)
	if err := cursor.All(ctx, &categories); err != nil {
		return make([]models.Category, 0), err
	}

	return categories, nil
}

func (r *CategoryRepository) FindById(ctx context.Context, idStr string) (models.Category, error) {
	objID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		slog.Debug("Error to convert id to object id", "error", err)
		return models.Category{}, err
	}

	filter := bson.M{"_id": objID}

	slog.Info("Searching category by id", "filter", filter)
	cursor := r.collection.FindOne(ctx, filter)

	var category models.Category
	if err := cursor.Decode(&category); err != nil {
		slog.Debug("Error to decode category", "error", err)
		return models.Category{}, err
	}

	return category, nil

}

func (r *CategoryRepository) CreateCategory(ctx context.Context, category models.Category) (*mongo.InsertOneResult, error) {
	slog.Debug("Creating category", "category", category)
	res, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		slog.Debug("Error to create category", "error", err)
		return nil, err
	}

	return res, nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, idStr string) (*mongo.DeleteResult, error) {
	objID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	slog.Debug("Deleting category by id", "filter", filter)
	res, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		slog.Debug("Error to delete category", "error", err)
		return nil, err
	}

	return res, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, idStr string, category models.Category) (*mongo.UpdateResult, error) {
	objID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		slog.Debug("Error to convert id to object id", "error", err)
		return nil, err
	}

	filter := bson.M{"_id": objID}

	update := bson.M{
		"$set": bson.M{
			"name":  category.Name,
			"type":  category.Type,
			"color": category.Color,
		},
	}

	slog.Debug("Updating category", "filter", filter, "update", update)
	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		slog.Debug("Error to update category", "error", err)
		return nil, err
	}
	slog.Debug("Category updated", "result", res)

	return res, nil
}
