package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupArticleNutritionRoutes(router *gin.Engine) {
	nutrition := router.Group("/ArticleNutrition")
	{
		nutrition.POST("/", handlers.CreateNutritionArticleHandler)
		nutrition.GET("", handlers.GetAllArticleNutritionsHandler)
		nutrition.GET("/:id", handlers.GetArticleNutritionByIDHandler)
		nutrition.PUT("/:id", handlers.UpdateNutritionArticleHandler)
		nutrition.DELETE("/:id", handlers.DeleteArticleNutritionHandler)
	}
}
