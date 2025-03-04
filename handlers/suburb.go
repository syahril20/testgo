package handlers

import (
	"context"
	"net/http"
	"server/dto/suburb"
	"server/models"
	"server/repositories"
	"time"

	resultdto "server/dto/result"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateSuburbHandler(c *gin.Context) {
	var suburbDTO dto.SuburbRequest
	if err := c.ShouldBindJSON(&suburbDTO); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Error in binding data: " + err.Error(),
		})
		return
	}

	cityID, err := primitive.ObjectIDFromHex(suburbDTO.CityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid CityID format: " + err.Error(),
		})
		return
	}

	suburb := models.Suburb{
		ID:        primitive.NewObjectID(),
		Name:      suburbDTO.Name,
		CityID:    cityID,
		CreatedAt: suburbDTO.CreatedAt,
		CreatedBy: suburbDTO.CreatedBy,
		UpdatedAt: suburbDTO.UpdatedAt,
		UpdatedBy: suburbDTO.UpdatedBy,
	}

	if err := repositories.CreateSuburb(context.Background(), suburb); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create suburb: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Suburb created successfully",
		Data:    suburb,
	})
}

func GetAllSuburbsHandler(c *gin.Context) {
	suburbs, err := repositories.GetAllSuburbs(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get suburbs: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Suburbs retrieved successfully",
		Data:    suburbs,
	})
}

func GetSuburbHandlerByID(c *gin.Context) {
	id := c.Param("id")
	suburb, err := repositories.GetSuburbByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Suburb not found"})
		return
	}
	c.JSON(http.StatusOK, suburb)
}

func UpdateSuburbHandler(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format: " + err.Error(),
		})
		return
	}

	oldSuburb, err := repositories.GetSuburbByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, resultdto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Suburb not found: " + err.Error(),
		})
		return
	}

	var suburbDTO dto.SuburbRequest
	if err := c.ShouldBindJSON(&suburbDTO); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Error in binding data: " + err.Error(),
		})
		return
	}

	cityID, err := primitive.ObjectIDFromHex(suburbDTO.CityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid CityID format: " + err.Error(),
		})
		return
	}

	suburb := models.Suburb{
		ID:        oldSuburb.ID,
		Name:      suburbDTO.Name,
		CityID:    cityID,
		CreatedAt: oldSuburb.CreatedAt,
		CreatedBy: oldSuburb.CreatedBy,
		UpdatedAt: time.Now(),
		UpdatedBy: suburbDTO.UpdatedBy,
	}

	if err := repositories.UpdateSuburb(context.Background(), objectID, suburb); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update suburb: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Suburb updated successfully",
		Data:    suburb,
	})
}

func DeleteSuburbHandler(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format: " + err.Error(),
		})
		return
	}

	suburb, err := repositories.GetSuburbByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, resultdto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Suburb not found: " + err.Error(),
		})
		return
	}

	deletedAt := time.Now()

	if err := repositories.DeleteSuburb(c, objectID, deletedAt); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete suburb: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Suburb deleted successfully",
		Data: gin.H{
			"ID":        suburb.ID.Hex(),
			"Name":      suburb.Name,
			"CityID":    suburb.CityID.Hex(),
			"CreatedAt": suburb.CreatedAt,
			"CreatedBy": suburb.CreatedBy,
			"UpdatedAt": suburb.UpdatedAt,
			"UpdatedBy": suburb.UpdatedBy,
			"DeletedAt": deletedAt,
		},
	})
}
