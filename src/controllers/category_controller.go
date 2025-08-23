package controllers

import (
	"context"
	"go-fcontrol-api/src/models"
	"go-fcontrol-api/src/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const defaultTimeout = 10 * time.Second

type CategoryController struct {
	service services.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		service: *services.NewCategoryService(),
	}
}

func (cc *CategoryController) GetCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var filter models.Category
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input binding",
			"message": err.Error(),
		})
		return
	}

	categories, err := cc.service.GetCategory(ctx, filter)
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

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input binding",
			"message": err.Error(),
		})
		return
	}

	res, err := cc.service.CreateCategory(ctx, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create category",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    res,
	})
}

func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input binding",
			"message": err.Error(),
		})
		return
	}

	res, err := cc.service.UpdateCategory(ctx, c.Param("id"), input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Error to update category",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    res,
	})
}

func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	id := c.Param("id")

	_, err := cc.service.DeleteCategory(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Error to delete category",
			"message": "Category not found",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
