package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(r *gin.Engine) {
	product := r.Group("/product")
	{
		product.POST("/", handlers.CreateProductHandler)
		product.GET("/", handlers.GetAllProductHandler)
	}

	subProduct := r.Group("/sub_product")
	{
		subProduct.POST("/", handlers.CreateSubProductHandler)
	}

	addons := r.Group("/addons")
	{
		addons.POST("/", handlers.CreateAddonsHandler)
	}
}
