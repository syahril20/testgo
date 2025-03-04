package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupArticleRoutes(r *gin.Engine) {
	article := r.Group("/articles")
	{
		article.POST("/", handlers.CreateArticleHandler)
		article.GET("/", handlers.GetAllArticlesHandler)
		article.GET("/:id", handlers.GetArticleHandler)
		article.PUT("/:id", handlers.UpdateArticleHandler)
		article.DELETE("/:id", handlers.DeleteArticleHandler)
	}
}
