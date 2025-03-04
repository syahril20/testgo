package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupFavouriteRoutes(r *gin.Engine) {
	favourite := r.Group("/favorites")
	{
		favourite.POST("/", handlers.CreateFavoriteHandler)
		favourite.GET("/", handlers.GetAllFavoritesHandler)
		favourite.GET("/:id", handlers.GetFavoriteByIDHandler)
		favourite.PUT("/:id", handlers.UpdateFavoriteHandler)
		favourite.DELETE("/:id", handlers.DeleteFavoriteHandler)
	}
}
