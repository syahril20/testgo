package handlers

import (
	"context"
	"net/http"
	"server/dto"
	"server/models"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateRingHandler(c *gin.Context) {
	var req dto.CreateRingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ring := models.Ring{
		Size:       req.Size,
		Color:      req.Color,
		Connection: req.Connection, // Menambahkan field Connection
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repositories.CreateRing(ctx, ring); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ring"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Ring created successfully"})
}

// GetRingHandler handles getting a ring by its ID
func GetRingHandler(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ring, err := repositories.GetRingByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ring not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ring": ring})
}

// UpdateRingHandler handles updating a ring by its ID
func UpdateRingHandler(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateRingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ring := models.Ring{
		Size:       req.Size,
		Color:      req.Color,
		Connection: req.Connection, // Menambahkan field Connection
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repositories.UpdateRing(ctx, id, ring); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ring"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ring updated successfully"})
}

// DeleteRingHandler handles deleting a ring by its ID
func DeleteRingHandler(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repositories.DeleteRing(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ring"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ring deleted successfully"})
}
