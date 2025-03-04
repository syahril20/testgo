package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/dto"
	"server/repositories"
)

// CreateFitnessCategory: Membuat kategori fitness
func CreateFitnessCategory(c *gin.Context) {
	var categoryDTO dto.FitnessCategoryDTO
	if err := c.ShouldBindJSON(&categoryDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.CreateFitnessCategory(context.Background(), categoryDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Fitness Category created successfully"})
}

// UpdateFitnessCategory: Memperbarui kategori fitness berdasarkan ID
func UpdateFitnessCategory(c *gin.Context) {
	id := c.Param("id")
	var categoryDTO dto.FitnessCategoryDTO
	if err := c.ShouldBindJSON(&categoryDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.UpdateFitnessCategory(context.Background(), id, categoryDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fitness Category updated successfully"})
}

// DeleteFitnessCategory: Menghapus kategori fitness berdasarkan ID
func DeleteFitnessCategory(c *gin.Context) {
	id := c.Param("id")

	if err := repositories.DeleteFitnessCategory(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fitness Category deleted successfully"})
}

// GetFitnessCategoryByID: Mengambil kategori fitness berdasarkan ID
func GetFitnessCategoryByID(c *gin.Context) {
	id := c.Param("id")
	category, err := repositories.GetFitnessCategoryByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fitness Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// GetAllFitnessCategories: Mengambil semua kategori fitness
func GetAllFitnessCategories(c *gin.Context) {
	categories, err := repositories.GetAllFitnessCategories(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
