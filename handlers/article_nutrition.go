package handlers

import (
	"context"
	"net/http"
	article_nutrition "server/dto/article/article_nutrition"
	resultdto "server/dto/result"
	"server/models"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNutritionArticleHandler - Menangani pembuatan artikel nutrisi baru
func CreateNutritionArticleHandler(c *gin.Context) {
	var request article_nutrition.NutritionArticleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var nutritions []models.Nutrition
	for _, nutrition := range request.Nutritions {
		nutritions = append(nutritions, models.Nutrition{
			Title: nutrition.Title,
			Value: nutrition.Value,
			Unit:  nutrition.Unit,
		})
	}

	article := models.NutritionArticle{
		ID:           primitive.NewObjectID(),
		CategoryID:   request.IDCategory,
		Title:        request.Title,
		Image:        request.Image,
		Content:      request.Content,
		TimeToCook:   request.TimeToCook,
		ServingSize:  request.ServingSize,
		Nutritions:   nutritions,
		Ingredients:  request.Ingredients,
		Instructions: request.Instructions,
		CreatedBy:    request.CreatedBy,
		CreatedAt:    time.Now(),
	}

	if err := repositories.CreateArticleNutrition(c.Request.Context(), article); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resultdto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "Nutrition article created successfully",
		Data:    article,
	})
}

// GetArticleNutritionByIDHandler - Mendapatkan artikel nutrisi berdasarkan ID
func GetArticleNutritionByIDHandler(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format: " + err.Error(),
		})
		return
	}

	article, err := repositories.GetArticleNutritionByID(context.Background(), objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, resultdto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Article nutrition not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Article nutrition retrieved successfully",
		Data:    article,
	})
}

func UpdateNutritionArticleHandler(c *gin.Context) {
	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var request article_nutrition.NutritionArticleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingArticle, err := repositories.GetArticleNutritionByID(c.Request.Context(), objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Update hanya field yang diperlukan
	article := models.NutritionArticle{
		ID:           objectID,
		Title:        request.Title,
		CategoryID:   request.IDCategory,
		Image:        request.Image,
		Content:      request.Content,
		Ingredients:  request.Ingredients,
		Instructions: request.Instructions,
		Nutritions:   convertToModelNutritions(request.Nutritions),
		UpdatedBy:    request.UpdatedBy,
		UpdatedAt:    time.Now(),
	}

	// Jangan hapus field yang tidak dikirimkan: gunakan nilai lama jika tidak ada
	if request.TimeToCook != "" {
		article.TimeToCook = request.TimeToCook
	} else {
		article.TimeToCook = existingArticle.TimeToCook
	}

	if request.ServingSize != "" {
		article.ServingSize = request.ServingSize
	} else {
		article.ServingSize = existingArticle.ServingSize
	}

	// Jangan ubah `CreatedBy` jika tidak disertakan dalam permintaan
	if request.CreatedBy != "" {
		article.CreatedBy = request.CreatedBy
	} else {
		article.CreatedBy = existingArticle.CreatedBy
	}

	// Pastikan `UpdatedAt` sudah di-set ke waktu saat ini
	article.UpdatedAt = time.Now()

	// Jangan ubah `CreatedAt`
	article.CreatedAt = existingArticle.CreatedAt

	if err := repositories.UpdateArticleNutrition(c.Request.Context(), objectID, article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Nutrition article updated successfully",
		"data":    article,
	})
}

func DeleteArticleNutritionHandler(c *gin.Context) {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Validasi format ID MongoDB
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Periksa apakah artikel nutrisi ada
	existingArticle, err := repositories.GetArticleNutritionByID(c.Request.Context(), objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Menambahkan timestamp pada DeletedAt untuk soft delete
	deletedAt := time.Now()
	existingArticle.DeletedAt = &deletedAt

	// Perbarui artikel dengan deletedAt
	if err := repositories.UpdateArticleNutrition(c.Request.Context(), objectID, existingArticle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update article with deletedAt: " + err.Error(),
		})
		return
	}

	// Kembalikan respon sukses dengan data artikel yang dihapus
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Article nutrition deleted successfully",
		"data":    existingArticle,
	})
}

// GetAllArticleNutritionsHandler - Mengambil semua artikel nutrisi
func GetAllArticleNutritionsHandler(c *gin.Context) {
	article, err := repositories.GetAllArticleNutritions(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch article nutritions: " + err.Error(),
		})
		return
	}

	if len(article) == 0 {
		c.JSON(http.StatusOK, resultdto.SuccessResult{
			Code:    http.StatusOK,
			Message: "No nutrition articles found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Article nutritions retrieved successfully",
		Data:    article,
	})
}

func convertToModelNutritions(dtoNutritions []article_nutrition.Nutrition) []models.Nutrition {
	var modelNutritions []models.Nutrition
	for _, n := range dtoNutritions {
		modelNutritions = append(modelNutritions, models.Nutrition{
			Title: n.Title,
			Value: n.Value,
			Unit:  n.Unit,
		})
	}
	return modelNutritions
}
