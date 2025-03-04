package handlers

import (
	"context"
	"net/http"
	"server/dto/article/article_category"
	"server/repositories"

	"github.com/gin-gonic/gin"
)

// CreateCategoryHandler: Menambahkan kategori baru
func CreateCategoryHandler(c *gin.Context) {
	var category dto.CategoryRequest
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repositories.CreateCategory(context.Background(), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

func GetCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	category, err := repositories.GetCategoryByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func UpdateCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	var category dto.CategoryRequest
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repositories.UpdateCategory(context.Background(), id, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

// DeleteCategoryHandler: Menghapus kategori berdasarkan ID
func DeleteCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	err := repositories.DeleteCategory(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// GetAllCategoriesHandler: Mengambil semua kategori
func GetAllCategoriesHandler(c *gin.Context) {
	categories, err := repositories.GetAllCategories(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}
