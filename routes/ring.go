package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRingRoutes(r *gin.Engine) {
	ring := r.Group("/rings")
	{
		ring.POST("/", handlers.CreateRingHandler)
		ring.GET("/:id", handlers.GetRingHandler)
		ring.PUT("/:id", handlers.UpdateRingHandler)
		ring.DELETE("/:id", handlers.DeleteRingHandler)
	}
}
