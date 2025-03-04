package handlers

import (
	"context"
	"fmt"
	"net/http"
	"server/dto"
	"server/repositories"

	"github.com/gin-gonic/gin"
)

func CreateFitnessActivity(c *gin.Context) {
	var activityDTO dto.FitnessActivityDTO

	// Bind JSON body ke activityDTO
	if err := c.ShouldBindJSON(&activityDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Panggil fungsi CreateFitnessActivity di repositori
	err := repositories.CreateFitnessActivity(c, activityDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create fitness activity: %v", err)})
		return
	}

	// Response sukses
	c.JSON(http.StatusCreated, gin.H{"message": "Fitness activity created successfully"})
}

func UpdateFitnessActivity(c *gin.Context) {
	id := c.Param("id")
	var activityDTO dto.FitnessActivityDTO
	if err := c.ShouldBindJSON(&activityDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.UpdateFitnessActivity(context.Background(), id, activityDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fitness Activity updated successfully"})
}

func DeleteFitnessActivity(c *gin.Context) {
	id := c.Param("id")

	if err := repositories.DeleteFitnessActivity(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fitness Activity deleted successfully"})
}

func GetFitnessActivityByID(c *gin.Context) {
	id := c.Param("id")
	activity, err := repositories.GetFitnessActivityByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fitness Activity not found"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func GetAllFitnessActivities(c *gin.Context) {
	activities, err := repositories.GetAllFitnessActivities(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activities)
}
