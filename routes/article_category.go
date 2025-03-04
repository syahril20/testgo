package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(r *gin.Engine) {
	category := r.Group("/categories")
	{
		category.POST("/", handlers.CreateCategoryHandler)
		category.GET("/:id", handlers.GetCategoryHandler)
		category.PUT("/:id", handlers.UpdateCategoryHandler)
		category.DELETE("/:id", handlers.DeleteCategoryHandler)
		category.GET("/", handlers.GetAllCategoriesHandler)
	}
}
