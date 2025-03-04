package handlers

import (
	"context"
	"net/http"
	"server/dto/city"
	resultdto "server/dto/result"
	"server/models"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCityHandler(c *gin.Context) {
	var cityDTO dto.CityRequest
	if err := c.ShouldBindJSON(&cityDTO); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Error in binding data: " + err.Error(),
		})
		return
	}

	provinceID, err := primitive.ObjectIDFromHex(cityDTO.ProvinceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Province ID format: " + err.Error(),
		})
		return
	}

	city := models.City{
		ID:         primitive.NewObjectID(),
		Name:       cityDTO.Name,
		ProvinceID: provinceID,
		CreatedAt:  time.Now(),
		CreatedBy:  cityDTO.CreatedBy,
		UpdatedAt:  time.Now(),
		UpdatedBy:  cityDTO.UpdatedBy,
	}

	if err := repositories.CreateCity(context.Background(), city); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create city: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "City created successfully",
		Data:    city,
	})
}

func GetCityHandlerByID(c *gin.Context) {
	id := c.Param("id")
	city, err := repositories.GetCityByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, resultdto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "City not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "City retrieved successfully",
		Data:    city,
	})
}

func UpdateCityHandler(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format: " + err.Error(),
		})
		return
	}

	oldCity, err := repositories.GetCityByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, resultdto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "City not found: " + err.Error(),
		})
		return
	}

	var cityDTO dto.CityRequest
	if err := c.ShouldBindJSON(&cityDTO); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Error in binding data: " + err.Error(),
		})
		return
	}

	provinceID, err := primitive.ObjectIDFromHex(cityDTO.ProvinceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Province ID format: " + err.Error(),
		})
		return
	}

	city := models.City{
		ID:         oldCity.ID,
		Name:       cityDTO.Name,
		ProvinceID: provinceID,
		CreatedAt:  oldCity.CreatedAt,
		CreatedBy:  oldCity.CreatedBy,
		UpdatedAt:  time.Now(),
		UpdatedBy:  cityDTO.UpdatedBy,
	}

	if err := repositories.UpdateCity(context.Background(), objectID, city); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update city: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "City updated successfully",
		Data:    city,
	})
}

func DeleteCityHandler(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format: " + err.Error(),
		})
		return
	}

	city, err := repositories.GetCityByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, resultdto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "City not found: " + err.Error(),
		})
		return
	}

	deletedAt := time.Now()

	err = repositories.SoftDeleteCity(context.Background(), objectID, deletedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to soft delete city: " + err.Error(),
		})
		return
	}

	city.DeletedAt = &deletedAt
	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "City deleted successfully",
		Data:    city,
	})
}

func GetAllCitiesHandler(c *gin.Context) {
	cities, err := repositories.GetAllCities(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve cities: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Cities retrieved successfully",
		Data:    cities,
	})
}
