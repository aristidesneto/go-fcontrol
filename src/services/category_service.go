package services

import (
	"fmt"
	"go-fcontrol-api/src/models"
	"go-fcontrol-api/src/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

func GetCategory(queryParam models.Category) ([]models.Category, error) {
	filter := bson.M{}

	if queryParam.Name != "" || queryParam.Type != "" {
		filter = bson.M{
			"$and": bson.A{
				bson.M{"name": bson.M{"$regex": queryParam.Name, "$options": "i"}},
				bson.M{"type": bson.M{"$regex": queryParam.Type, "$options": "i"}},
			},
		}
	}

	return repositories.GetCategory(filter)
}

func CreateCategory(category models.Category) (*models.Category, error) {
	res, err := repositories.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	// category.ID = res.InsertedID.(primitive.ObjectID)
	fmt.Println(res)

	return &category, nil
}
