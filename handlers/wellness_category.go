package handlers

import (
	"context"
	"net/http"
	dto "server/dto/result"
	WellnessCategoryDto "server/dto/wellness_category"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateWellnessCategoryHandler(c *gin.Context) {
	var req WellnessCategoryDto.CreateWellnessCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Content == "" || req.Title == "" || req.Image == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	WellnessCategory, _ := repositories.GetWellnessCategoryByName(context.Background(), req.Title)
	if WellnessCategory != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "gagal"})
		return
	}

	currentTime := time.Now()

	WellnessCategorys := WellnessCategoryDto.CreateWellnessCategoryRequest{
		Title:     req.Title,
		Content:   req.Content,
		Image:     req.Image,
		CreatedAt: currentTime,
		CreatedBy: "SYSTEM",
		UpdatedAt: currentTime,
		UpdatedBy: "SYSTEM",
	}

	// Menambahkan Suburb ke database
	data, err := repositories.CreateWellnessCategory(context.Background(), WellnessCategorys)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    convertResponseWellnessCategory(data)})
}

func convertResponseWellnessCategory(wellnessCategory WellnessCategoryDto.CreateWellnessCategoryRequest) WellnessCategoryDto.WellnessCategoryResponse {
	return WellnessCategoryDto.WellnessCategoryResponse{
		Title:   wellnessCategory.Title,
		Content: wellnessCategory.Content,
		Image:   wellnessCategory.Image,
	}
}
