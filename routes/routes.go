package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	SetupUserRoutes(r)
	SetupRingRoutes(r)
	SetupAddressRoutes(r)
	SetupCategoryRoutes(r)
	SetupArticleRoutes(r)
	SetupFitnessRoutes(r)
	SetupSuburbRoutes(r)
	SetupProductRoutes(r)
	SetupWellnessCategoryRoutes(r)
	SetupArticleNutritionRoutes(r)
	SetupassessmentRoutes(r)
	SetupPayment(r)
	SetupInvoice(r)
	SetupHpiRoutes(r)
	SetupMemberRoutes(r)
	SetupMembershipRoutes(r)
	SetupIMSRoutes(r)
	SetupFavouriteRoutes(r)
	SetupGamificationRoutes(r)
	SetupAtuhRoutes(r)
}
