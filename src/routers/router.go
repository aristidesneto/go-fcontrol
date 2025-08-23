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

	// expense := router.Group("/expense")
	// {
	// 	expense.GET("/", controllers.GetExpense)
	// 	expense.POST("/", controllers.CreateExpense)
	// }

	// user := router.Group("/user")
	// {
	// 	user.GET("/", controllers.GetUser)
	// 	user.POST("/", controllers.StoreUser)
	// }

	categoryController := controllers.NewCategoryController()

	category := router.Group("/category")
	{
		category.GET("/all", categoryController.GetCategory)
		category.POST("", categoryController.CreateCategory)
		// category.PUT("/:id", controllers.UpdateCategory)
		category.DELETE("/:id", categoryController.DeleteCategory)
	}
}
