package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.POST("/", handlers.CreateUserHandler)
		user.GET("/:id", handlers.GetUserByIDHandler)
		user.GET("/", handlers.GetUserHandler)
		user.PUT("/:id", handlers.UpdateUserHandler)
		user.PUT("/password/:id", handlers.UpgradePasswordHandler)
	}
}
