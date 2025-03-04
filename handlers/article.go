package handlers

import (
	"net/http"
	"server/dto/article"
	"server/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateArticleHandler(c *gin.Context) {
	var article dto.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	// Mengonversi idCategory dari string ke ObjectID
	idCategory, err := primitive.ObjectIDFromHex(article.IDCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid category ID format",
			"error":   err.Error(),
		})
		return
	}

	// Mengonversi kembali ObjectID ke string jika perlu untuk response
	article.IDCategory = idCategory.Hex()

	// Memanggil repository untuk menyimpan artikel
	err = repositories.CreateArticle(c.Request.Context(), article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create article",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Article created successfully",
		"data":    article,
	})
}

func GetArticleHandler(c *gin.Context) {
	id := c.Param("id")

	article, err := repositories.GetArticleByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Article not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   article,
	})
}

func UpdateArticleHandler(c *gin.Context) {
	id := c.Param("id")

	var article dto.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	err := repositories.UpdateArticle(c.Request.Context(), id, article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to update article",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Article updated successfully",
	})
}

func DeleteArticleHandler(c *gin.Context) {
	id := c.Param("id")

	err := repositories.DeleteArticle(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to delete article",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Article deleted successfully",
	})
}

func GetAllArticlesHandler(c *gin.Context) {
	articles, err := repositories.GetAllArticles(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch articles",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   articles,
	})
}
