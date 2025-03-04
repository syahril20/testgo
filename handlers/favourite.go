package handlers

import (
	"context"
	"net/http"
	"server/dto/favourite"
	"server/dto/result"
	"server/models"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFavoriteHandler(c *gin.Context) {
	var favoriteDTO dto.FavouriteDTO
	if err := c.ShouldBindJSON(&favoriteDTO); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	userID, _ := primitive.ObjectIDFromHex(favoriteDTO.UserID)
	articleID, _ := primitive.ObjectIDFromHex(favoriteDTO.ArticleID)

	favorite := models.Favourite{
		ID:        primitive.NewObjectID(),
		UserID:    userID,
		ArticleID: articleID,
		Count:     favoriteDTO.Count,
		CreatedAt: time.Now(),
		CreatedBy: favoriteDTO.CreatedBy,
	}

	err := repositories.CreateFavorite(context.Background(), favorite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Message: "Favorite created successfully", Data: favorite})
}

func GetFavoriteByIDHandler(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
		})
		return
	}

	favourite, err := repositories.GetFavoriteByID(c.Request.Context(), objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve Favourite: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Favourite successfully",
		Data:    favourite,
	})
}

func GetAllFavoritesHandler(c *gin.Context) {
	favorites, err := repositories.GetAllFavorites(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Message: "Favorites retrieved successfully", Data: favorites})
}

func UpdateFavoriteHandler(c *gin.Context) {
	id := c.Param("id")
	var favoriteDTO dto.FavouriteUpdateDTO // Pastikan nama di sini adalah favoriteDTO

	if err := c.ShouldBindJSON(&favoriteDTO); err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid input: " + err.Error(),
		})
		return
	}

	// Validasi ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
		})
		return
	}

	// Update di database
	err = repositories.UpdateFavorite(c.Request.Context(), objectID, favoriteDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update favorite: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Favorite updated successfully",
		Data:    favoriteDTO,
	})
}

func DeleteFavoriteHandler(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
		})
		return
	}

	// Soft delete dengan mengubah field DeletedAt
	favourite, err := repositories.GetFavoriteByID(c.Request.Context(), objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to soft delete favorite: " + err.Error(),
		})
		return
	}

	now := time.Now()
	if err := repositories.DeleteFavorite(c.Request.Context(), objectID, now); err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete Favourite: " + err.Error(),
		})
		return
	}

	favourite.DeletedAt = now
	c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:    http.StatusOK,
		Message: "IMS deleted successfully (soft delete)",
		Data:    favourite,
	})
}
