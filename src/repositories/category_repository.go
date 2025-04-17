package repositories

import (
	"context"
	"go-fcontrol-api/src/configs"
	"go-fcontrol-api/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var categoryCollection = configs.GetCollection(configs.DB, "categories")

func GetCategory(filter bson.M) ([]models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := categoryCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []models.Category
	if err := cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func CreateCategory(category models.Category) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := categoryCollection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}

	return res, nil
}
