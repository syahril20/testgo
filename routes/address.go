package routes

import (
	"github.com/gin-gonic/gin"
	"server/handlers"
)

func SetupAddressRoutes(r *gin.Engine) {
	address := r.Group("/address")
	{
		// Province routes
		address.POST("/province", handlers.CreateProvinceHandler)
		address.GET("/province", handlers.GetAllProvincesHandler)
		address.GET("/province/:id", handlers.GetProvinceHandlerByID)
		address.PUT("/province/:id", handlers.UpdateProvinceHandler)
		address.DELETE("/province/:id", handlers.DeleteProvinceHandler)

		// City routes
		address.POST("/city", handlers.CreateCityHandler)
		address.GET("/city", handlers.GetAllCitiesHandler)
		address.GET("/city/:id", handlers.GetCityHandlerByID)
		address.PUT("/city/:id", handlers.UpdateCityHandler)
		address.DELETE("/city/:id", handlers.DeleteCityHandler)

		// Suburb routes
		address.POST("/suburb", handlers.CreateSuburbHandler)
		address.GET("/suburb", handlers.GetAllSuburbsHandler)
		address.GET("/suburb/:id", handlers.GetSuburbHandlerByID)
		address.PUT("/suburb/:id", handlers.UpdateSuburbHandler)
		address.DELETE("/suburb/:id", handlers.DeleteSuburbHandler)

		// Area routes
		address.POST("/area", handlers.CreateAreaHandler)
		address.GET("/area", handlers.GetAllAreaHandler)
		address.GET("/area/:id", handlers.GetAreaHandlerByID)
		address.PUT("/area/:id", handlers.UpdateAreaHandler)
		address.DELETE("/area/:id", handlers.SoftDeleteAreaHandler)
	}
}
