package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupHpiRoutes(r *gin.Engine) {
	hpi := r.Group("/hpi")
	{
		hpi.GET("/", handlers.GetAllActiveHpiHandler)
		hpi.GET("/:id", handlers.GetHpiHandlerById)
		hpi.GET("/non-active", handlers.GetAllNonActiveHpiHandler)
		hpi.POST("/", handlers.CreateHpiHandler)
		hpi.PUT("/:id", handlers.UpdateHpiHandler)
		hpi.PUT("/active/:id", handlers.ActiveHpiHandler)
		hpi.DELETE("/:id", handlers.DeleteHpiHandler)
	}

	biomarker := r.Group("/biomarker")
	{
		biomarker.GET("/:id", handlers.GetAllActiveBiomarkersHandler)
		biomarker.GET("/", handlers.GetBiomarkerHandlerById)
		biomarker.GET("/non-active/:id", handlers.GetAllNonActiveBiomarkersHandler)
		biomarker.POST("/", handlers.CreateBiomarkerHandler)
		biomarker.PUT("/:id", handlers.UpdateBiomarkerHandler)
		biomarker.PUT("/active/:id", handlers.ActiveBiomarkerHandler)
		biomarker.DELETE("/:id", handlers.DeleteBiomarkerHandler)
	}

	under := r.Group("/under")
	{
		under.GET("/:id", handlers.GetUnderByBiomarkerIdHandler)
		under.POST("/:id", handlers.CreateUnderHandler)
		under.PUT("/:id", handlers.UpdateUnderHandler)
	}

	over := r.Group("/over")
	{
		over.GET("/:id", handlers.GetOverByBiomarkerIdHandler)
		over.POST("/:id", handlers.CreateOverHandler)
		over.PUT("/:id", handlers.UpdateOverHandler)
	}

	lifestyle := r.Group("/lifestyle")
	{
		lifestyle.PUT("/:id", handlers.UpdateLifestyleHandler)
	}

	result := r.Group("/hpi_result")
	{
		result.GET("/:id", handlers.GetHpiResultHandlerById)
		result.POST("/", handlers.CreateHpiResultsHandler)
	}
}
