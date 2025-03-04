package handlers

import (
	"net/http"
	"server/dto"
	"server/repositories"

	"github.com/gin-gonic/gin"
)

func CreateFitnessHandler(c *gin.Context) {
	var fitness dto.FitnessDTO
	if err := c.ShouldBindJSON(&fitness); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	// Memanggil CreateFitness dan menambahkan ID
	createdFitness, err := repositories.CreateFitness(c.Request.Context(), fitness)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create fitness",
			"error":   err.Error(),
		})
		return
	}

	// Mengembalikan response dengan fitness yang sudah lengkap
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Fitness created successfully",
		"data":    createdFitness,
	})
}

// GetFitnessHandler: Mendapatkan fitness berdasarkan ID
func GetFitnessHandler(c *gin.Context) {
	id := c.Param("id")

	fitness, err := repositories.GetFitnessByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Fitness not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   fitness,
	})
}

// GetAllFitnessHandler: Mendapatkan semua data fitness
func GetAllFitnessHandler(c *gin.Context) {
	fitnessList, err := repositories.GetAllFitness(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch fitness",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   fitnessList,
	})
}

// UpdateFitnessHandler: Memperbarui fitness berdasarkan ID
func UpdateFitnessHandler(c *gin.Context) {
	id := c.Param("id")

	var fitness dto.FitnessDTO
	if err := c.ShouldBindJSON(&fitness); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	err := repositories.UpdateFitness(c.Request.Context(), id, fitness)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to update fitness",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Fitness updated successfully",
	})
}

// DeleteFitnessHandler: Menghapus fitness berdasarkan ID
func DeleteFitnessHandler(c *gin.Context) {
	id := c.Param("id")

	err := repositories.DeleteFitness(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to delete fitness",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Fitness deleted successfully",
	})
}
