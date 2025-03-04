package handlers

import (
	"net/http"
	"server/dto"
	resultdto "server/dto/result"
	"server/models"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFitnessProgram(c *gin.Context) {
	var request dto.CreateFitnessProgramRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	program := models.FitnessProgram{
		ID:          primitive.NewObjectID(),
		Title:       request.Title,
		Image:       request.Image,
		Description: request.Description,
		Instructor:  request.Instructor,
		Date:        time.Now(),
		Trait:       request.Trait,
		Duration:    request.Duration,
		Category:    request.Category,
		Link:        request.Link,
		CreatedAt:   time.Now(),
		CreatedBy:   request.CreatedBy,
	}

	if err := repositories.CreateFitnessProgram(c.Request.Context(), program); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resultdto.SuccessResult{Code: http.StatusCreated, Message: "Fitness Program created successfully", Data: program})
}

func GetAllFitnessPrograms(c *gin.Context) {
	programs, err := repositories.GetAllFitnessPrograms(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: programs})
}

func GetFitnessProgramByID(c *gin.Context) {
	id := c.Param("id")

	program, err := repositories.GetFitnessProgramByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, resultdto.ErrorResult{Code: http.StatusNotFound, Message: "Fitness Program not found"})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: program})
}

func UpdateFitnessProgram(c *gin.Context) {
	id := c.Param("id")
	var request dto.UpdateFitnessProgramRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if request.Title != "" {
		updates["title"] = request.Title
	}
	if request.Image != "" {
		updates["image"] = request.Image
	}
	if request.Description != "" {
		updates["description"] = request.Description
	}
	if request.Instructor != "" {
		updates["instructor"] = request.Instructor
	}
	if len(request.Trait) > 0 {
		updates["trait"] = request.Trait
	}
	if request.Duration != "" {
		updates["duration"] = request.Duration
	}
	if request.Category != "" {
		updates["category"] = request.Category
	}
	if request.Link != "" {
		updates["link"] = request.Link
	}

	if err := repositories.UpdateFitnessProgram(c.Request.Context(), id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Message: "Fitness Program updated successfully"})
}

func DeleteFitnessProgram(c *gin.Context) {
	id := c.Param("id")

	if err := repositories.DeleteFitnessProgram(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Message: "Fitness Program deleted successfully"})
}
