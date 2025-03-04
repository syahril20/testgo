package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupMemberRoutes(r *gin.Engine) {
	member := r.Group("/member")
	{
		member.GET("/", handlers.GetAllMembersHandler)
		member.GET("/non-active", handlers.GetNonActiveAllMembersHandler)
		member.GET("/:id", handlers.GetMemberByIDHandler)
		member.POST("/", handlers.CreateMemberHandler)
		member.PUT("/:id", handlers.UpdateMemberByIDHandler)
		member.PUT("/active/:id", handlers.ActiveMemberHandler)
		member.DELETE("/:id", handlers.DeleteMemberHandler)
	}
}
