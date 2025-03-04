package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupInvoice(r *gin.Engine) {
	invoice := r.Group("/invoices")
	{
		invoice.POST("", handlers.CreateInvoiceHandler)
		invoice.GET("/:id", handlers.GetInvoiceByIDHandler)
		invoice.PUT("/:id", handlers.UpdateInvoiceHandler)
		invoice.DELETE("/:id", handlers.SoftDeleteInvoiceHandler)
	}
}
