package controllers

import (
	"context"
	"fmt"
	"go-fcontrol-api/src/configs"
	"go-fcontrol-api/src/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var expenseCollection = configs.GetCollection(configs.DB, "expenses")

func GetExpense(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := expenseCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer cursor.Close(ctx)

	var expenses []models.Expense
	for cursor.Next(ctx) {
		fmt.Println(cursor.Current)
		var expense models.Expense
		if err := cursor.Decode(&expense); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		expenses = append(expenses, expense)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    expenses,
	})
}

func CreateExpense(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input models.Expense
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	expense := models.Expense{
		Name:       input.Name,
		CategoryID: input.CategoryID,
	}

	res, err := expenseCollection.InsertOne(ctx, expense)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    res.InsertedID,
	})
}
