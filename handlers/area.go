package handlers

import (
	"context"
	"net/http"
	"server/dto/result"
	"server/models"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAreaHandler(c *gin.Context) {
	var area models.Area
	if err := c.ShouldBindJSON(&area); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid input data: " + err.Error(),
		})
		return
	}

	area.ID = primitive.NewObjectID()
	area.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	area.UpdatedAt = area.CreatedAt

	if err := repositories.CreateArea(context.Background(), area); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create area: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Area created successfully",
		Data:    area,
	})
}

func GetAllAreaHandler(c *gin.Context) {
	areas, err := repositories.GetAllAreas(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve areas: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Areas retrieved successfully",
		Data:    areas,
	})
}

func GetAreaHandlerByID(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format: " + err.Error(),
		})
		return
	}

	area, err := repositories.GetAreaByID(context.Background(), objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, resultdto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Area not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Area retrieved successfully",
		Data:    area,
	})
}

func UpdateAreaHandler(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format: " + err.Error(),
		})
		return
	}

	oldArea, err := repositories.GetAreaByID(context.Background(), objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Area not found: " + err.Error(),
		})
		return
	}

	var updatedArea models.Area
	if err := c.ShouldBindJSON(&updatedArea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	updatedArea.ID = oldArea.ID
	updatedArea.CreatedAt = oldArea.CreatedAt
	updatedArea.CreatedBy = oldArea.CreatedBy
	updatedArea.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	if updatedArea.UpdatedBy == "" {
		updatedArea.UpdatedBy = "system"
	}

	if err := repositories.UpdateArea(context.Background(), objectID, updatedArea); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update area: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Area updated successfully",
		"data":    updatedArea,
	})
}

func SoftDeleteAreaHandler(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format: " + err.Error(),
		})
		return
	}

	area, err := repositories.GetAreaByID(context.Background(), objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Area not found: " + err.Error(),
		})
		return
	}

	deletedAt := primitive.NewDateTimeFromTime(time.Now())
	area.DeletedAt = &deletedAt

	if err := repositories.SoftDeleteArea(context.Background(), objectID, deletedAt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete area: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Area deleted successfully",
		"data":    area,
	})
}
