package handlers

import (
	"context"
	"net/http"
	"server/dto/province"
	resultdto "server/dto/result"
	"server/models"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProvinceHandler(c *gin.Context) {
	var provinceDTO dto.ProvinceRequest
	if err := c.ShouldBindJSON(&provinceDTO); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Error in binding data: " + err.Error(),
		})
		return
	}

	province := models.Province{
		ID:        primitive.NewObjectID(),
		Name:      provinceDTO.Name,
		CreatedAt: provinceDTO.CreatedAt,
		CreatedBy: provinceDTO.CreatedBy,
		UpdatedAt: provinceDTO.UpdatedAt,
		UpdatedBy: provinceDTO.UpdatedBy,
	}

	if err := repositories.CreateProvince(context.Background(), province); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create province: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Province created successfully",
		Data:    province,
	})
}

func GetProvinceHandlerByID(c *gin.Context) {
	id := c.Param("id")
	province, err := repositories.GetProvinceByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, resultdto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Province not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Province retrieved successfully",
		Data:    province,
	})
}

func UpdateProvinceHandler(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format: " + err.Error(),
		})
		return
	}

	oldProvince, err := repositories.GetProvinceByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, resultdto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Province not found: " + err.Error(),
		})
		return
	}

	var provinceDTO dto.ProvinceRequest
	if err := c.ShouldBindJSON(&provinceDTO); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Error in binding data: " + err.Error(),
		})
		return
	}

	province := models.Province{
		ID:        oldProvince.ID,
		Name:      provinceDTO.Name,
		CreatedAt: oldProvince.CreatedAt,
		CreatedBy: oldProvince.CreatedBy,
		UpdatedAt: time.Now(),
		UpdatedBy: provinceDTO.UpdatedBy,
	}

	if err := repositories.UpdateProvince(context.Background(), objectID, province); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update province: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Province updated successfully",
		Data:    province,
	})
}

func DeleteProvinceHandler(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format: " + err.Error(),
		})
		return
	}

	province, err := repositories.GetProvinceByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, resultdto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Province not found: " + err.Error(),
		})
		return
	}

	deletedAt := time.Now()

	err = repositories.SoftDeleteProvince(context.Background(), objectID, deletedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to soft delete province: " + err.Error(),
		})
		return
	}

	province.DeletedAt = &deletedAt
	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Province deleted successfully",
		Data:    province,
	})
}

func GetAllProvincesHandler(c *gin.Context) {
	provinces, err := repositories.GetAllProvinces(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve provinces: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Provinces retrieved successfully",
		Data:    provinces,
	})
}
