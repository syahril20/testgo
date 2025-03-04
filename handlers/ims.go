package handlers

import (
	"net/http"
	"server/dto"
	result "server/dto/result"
	"server/models"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateIMSHandler(c *gin.Context) {
	var imsDTO dto.IMSRequest
	if err := c.ShouldBindJSON(&imsDTO); err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid input: " + err.Error(),
		})
		return
	}

	ims := models.IMS{
		ID:                   primitive.NewObjectID(),
		Email:                imsDTO.Email,
		Name:                 imsDTO.Name,
		Old:                  imsDTO.Old,
		Phone:                imsDTO.Phone,
		Address:              imsDTO.Address,
		OptiSampleCollection: imsDTO.OptiSampleCollection,
		CreatedAt:            time.Now(),
		CreatedBy:            imsDTO.CreatedBy,
	}

	if err := repositories.CreateIMS(c.Request.Context(), ims); err != nil {
		c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create IMS: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result.SuccessResult{
		Code:    http.StatusCreated,
		Message: "IMS created successfully",
		Data:    ims,
	})
}

func GetAllIMSHandler(c *gin.Context) {
	imsList, err := repositories.GetAllIMS(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve IMS data: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result.SuccessResult{
		Code:    http.StatusOK,
		Message: "IMS data retrieved successfully",
		Data:    imsList,
	})
}

func GetIMSByIDHandler(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
		})
		return
	}

	ims, err := repositories.GetIMSByID(c.Request.Context(), objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve IMS: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result.SuccessResult{
		Code:    http.StatusOK,
		Message: "IMS retrieved successfully",
		Data:    ims,
	})
}

func UpdateIMSHandler(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
		})
		return
	}

	var imsDTO dto.IMSRequest
	if err := c.ShouldBindJSON(&imsDTO); err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid input: " + err.Error(),
		})
		return
	}

	if err := repositories.UpdateIMS(c.Request.Context(), objectID, imsDTO); err != nil {
		c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update IMS: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result.SuccessResult{
		Code:    http.StatusOK,
		Message: "IMS updated successfully",
	})
}

func DeleteIMSHandler(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
		})
		return
	}

	// Ambil data IMS sebelum soft delete
	ims, err := repositories.GetIMSByID(c.Request.Context(), objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, result.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "IMS not found",
		})
		return
	}

	// Lakukan soft delete (set deletedAt)
	now := time.Now()
	if err := repositories.DeleteIMS(c.Request.Context(), objectID, now); err != nil {
		c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete IMS: " + err.Error(),
		})
		return
	}

	// Kembalikan data IMS yang dihapus
	ims.DeletedAt = now
	c.JSON(http.StatusOK, result.SuccessResult{
		Code:    http.StatusOK,
		Message: "IMS deleted successfully (soft delete)",
		Data:    ims,
	})
}
