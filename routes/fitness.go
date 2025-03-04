package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupFitnessRoutes(r *gin.Engine) {
	fitness := r.Group("/fitness")
	{
		fitness.POST("/", handlers.CreateFitnessHandler)
		fitness.GET("/", handlers.GetAllFitnessHandler)
		fitness.GET("/:id", handlers.GetFitnessHandler)
		fitness.PUT("/:id", handlers.UpdateFitnessHandler)
		fitness.DELETE("/:id", handlers.DeleteFitnessHandler)
	}

	fitnessCategory := r.Group("/fitness_category")
	{
		fitnessCategory.POST("/", handlers.CreateFitnessCategory)
		fitnessCategory.PUT("/:id", handlers.UpdateFitnessCategory)
		fitnessCategory.DELETE("/:id", handlers.DeleteFitnessCategory)
		fitnessCategory.GET("/:id", handlers.GetFitnessCategoryByID)
		fitnessCategory.GET("/", handlers.GetAllFitnessCategories)
	}
	fitnessActivity := r.Group("/fitness-activities")
	{
		fitnessActivity.POST("/", handlers.CreateFitnessActivity)
		fitnessActivity.GET("/", handlers.GetAllFitnessActivities)
		fitnessActivity.GET("/:id", handlers.GetFitnessActivityByID)
		fitnessActivity.PUT("/:id", handlers.UpdateFitnessActivity)
		fitnessActivity.DELETE("/:id", handlers.DeleteFitnessActivity)
	}

	fitnessProgram := r.Group("/fitness-programs")
	{
		fitnessProgram.POST("", handlers.CreateFitnessProgram)
		fitnessProgram.GET("", handlers.GetAllFitnessPrograms)
		fitnessProgram.GET("/:id", handlers.GetFitnessProgramByID)
		fitnessProgram.PUT("/:id", handlers.UpdateFitnessProgram)
		fitnessProgram.DELETE("/:id", handlers.DeleteFitnessProgram)
	}
}
