package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupWellnessCategoryRoutes(r *gin.Engine) {
	user := r.Group("/wellnessCategory")
	{
		user.POST("/", handlers.CreateWellnessCategoryHandler)
	}
}
