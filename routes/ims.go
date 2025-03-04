package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupIMSRoutes(r *gin.Engine) {
	ims := r.Group("/ims")
	ims.GET("", handlers.GetAllIMSHandler)
	ims.GET("/:id", handlers.GetIMSByIDHandler)
	ims.POST("", handlers.CreateIMSHandler)
	ims.PUT("/:id", handlers.UpdateIMSHandler)
	ims.DELETE("/:id", handlers.DeleteIMSHandler)
}
