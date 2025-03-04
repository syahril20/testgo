package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupGamificationRoutes(r *gin.Engine) {
	gamification := r.Group("/gamification")
	{
		gamification.GET("/", handlers.GetAllGamificationHandler)
		gamification.POST("/", handlers.CreateGamificationRequestHandler)
		gamification.PUT("/", handlers.UpdateGamificationHandler)
		gamification.PUT("/active", handlers.ActiveGamificationHandler)
		gamification.DELETE("/", handlers.DeleteGamificationHandler)
	}

	challenges := r.Group("/challenges")
	{
		challenges.GET("/", handlers.GetAllChallenges)
		challenges.POST("/", handlers.AddChallengesHandler)
		challenges.PUT("/:id", handlers.UpdateChallengesByIdHandler)
	}
}
