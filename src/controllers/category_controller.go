package controllers

import (
	"context"
	"go-fcontrol-api/src/configs"
	"go-fcontrol-api/src/models"
	"go-fcontrol-api/src/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var categoryCollection *mongo.Collection = configs.GetCollection(configs.DB, "categories")

func GetCategory(c *gin.Context) {
	var filter models.Category
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input binding",
			"message": err.Error(),
		})
		return
	}

	categories, err := services.GetCategory(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    categories,
	})
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input binding",
			"message": err.Error(),
		})
		return
	}

	res, err := services.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    res,
	})
}

func UpdateCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input binding",
			"message": err.Error(),
		})
		return
	}

	category, err := findById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Category not found",
			"message": err.Error(),
		})
		return
	}

	filter := bson.M{"_id": category.ID}
	update := bson.M{"$set": bson.M{
		"name": input.Name,
		"type": input.Type,
	}}

	res, err := categoryCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update category",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    res.ModifiedCount,
	})

}

func DeleteCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	category, err := findById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Category not found",
			"message": err.Error(),
		})
		return
	}

	filter := bson.M{"_id": category.ID}
	res, err := categoryCollection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Failed to delete category",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    res.DeletedCount,
	})
}

func findById(idHex string) (models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := bson.ObjectIDFromHex(idHex)
	if err != nil {
		return models.Category{}, err
	}

	filter := bson.M{"_id": id}

	var category models.Category
	err = categoryCollection.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}
