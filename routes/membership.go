package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupMembershipRoutes(r *gin.Engine) {
	membership := r.Group("/membership")
	{
		membership.GET("/", handlers.GetAllMembershipHandler)
		membership.POST("/", handlers.CreateMembershipHandler)
		membership.PUT("/:id", handlers.UpdateMembershipByIdHandler)
		membership.PUT("/active/:id", handlers.ActiveMembershipByIdHandler)
		membership.DELETE("/:id", handlers.DeleteMembershipHandler)
	}
}
