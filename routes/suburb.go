package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupSuburbRoutes(r *gin.Engine) {
	suburb := r.Group("/suburb")
	{
		suburb.POST("/", handlers.CreateSuburbHandler)
		suburb.GET("/:email", handlers.GetSuburbHandlerByID)
		suburb.PUT("/:email", handlers.UpdateUserHandler)
		suburb.PUT("/:email/password", handlers.UpgradePasswordHandler)
	}
}
