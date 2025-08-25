package routers

import (
	"go-fcontrol-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to FControl API 2",
		})
	})

	// Transactions
	transactionController := controllers.NewTransactionController()
	expense := router.Group("/transaction")
	{
		expense.GET("", transactionController.GetTransaction)
		expense.POST("", transactionController.CreateTransaction)
	}

	// Users
	userController := controllers.NewUserController()
	user := router.Group("/user")
	{
		user.GET("", userController.GetUser)
		user.POST("", userController.CreateUser)
	}

	// Categories
	categoryController := controllers.NewCategoryController()
	category := router.Group("/category")
	{
		category.GET("/all", categoryController.GetCategory)
		category.POST("", categoryController.CreateCategory)
		category.PUT("/:id", categoryController.UpdateCategory)
		category.DELETE("/:id", categoryController.DeleteCategory)
	}
}
