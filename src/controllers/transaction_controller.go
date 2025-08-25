package controllers

import (
	"context"
	"go-fcontrol-api/src/models"
	"go-fcontrol-api/src/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type TransactionController struct {
	service services.TransactionService
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		service: *services.NewTransactionService(),
	}
}

func (cc *TransactionController) GetTransaction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var filter models.Transaction
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input binding",
			"message": err.Error(),
		})
		return
	}

	transactions, err := cc.service.GetTransactions(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    transactions,
	})
}

func (cc *TransactionController) CreateTransaction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON payload",
			"details": err.Error(),
		})
		return
	}

	// Validação
	if err := validate.Struct(input); err != nil {
		var errors []map[string]string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, map[string]string{
				"field":   e.Field(),
				"rule":    e.Tag(),
				"message": validationMessage(e),
			})
		}

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":  "Validation failed",
			"fields": errors,
		})
		return
	}

	res, err := cc.service.CreateTransaction(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create transaction",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Transaction created successfully",
		"data":    res,
	})
}

func validationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "gt":
		return e.Field() + " must be greater than " + e.Param()
	case "datetime":
		return e.Field() + " must be a valid date (format: YYYY-MM-DD)"
	default:
		return e.Field() + " is invalid"
	}
}
