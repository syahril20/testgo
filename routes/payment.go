package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupPayment(r *gin.Engine) {
	payment := r.Group("/payment")
	{
		payment.POST("", handlers.CreatePaymentHandler)
		payment.GET("/:id", handlers.GetPaymentHandler)
		payment.PUT("/:id", handlers.UpdatePaymentHandler)
		payment.DELETE("/:id", handlers.DeletePaymentHandler)
	}
}
