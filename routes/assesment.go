package routes

import (
	"server/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupassessmentRoutes(r *gin.Engine) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Bisa diganti dengan domain frontend Anda
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	assessment := r.Group("/assessment")
	{
		assessment.GET("/", handlers.GetAllActiveassessmentsHandler)
		assessment.GET("/non-active", handlers.GetAllNonActiveassessmentsHandler)
		assessment.POST("/", handlers.CreateassessmentHandler)
		assessment.PUT("/:id", handlers.UpdateassessmentHandler)
		assessment.PUT("/active/:id", handlers.ActiveassessmentHandler)
		assessment.DELETE("/:id", handlers.DeleteassessmentHandler)
	}

	Questionnaire := r.Group("/questionnaire")
	{
		Questionnaire.GET("/:id", handlers.GetAllActiveQuestionnaireHandler)
		Questionnaire.GET("/non-active/:id", handlers.GetAllNonActiveQuestionnaireHandler)
		Questionnaire.POST("/", handlers.CreateQuestionnaireHandler)
		Questionnaire.PUT("/:id", handlers.UpdateQuestionnaireHandler)
		Questionnaire.PUT("/active/:id", handlers.ActiveQuestionnaireHandler)
		Questionnaire.DELETE("/:id", handlers.DeleteQuestionnaireHandler)
	}

	QuestionnaireResult := r.Group("/questionnaire_result")
	{
		QuestionnaireResult.POST("/", handlers.GetAssessmentPayloadHandler)
	}
}
