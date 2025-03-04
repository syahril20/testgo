package handlers

import (
	"context"
	"net/http"
	dtoGamification "server/dto/gamification"
	dto "server/dto/result"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateGamificationRequestHandler(c *gin.Context) {
	var req dtoGamification.CreateGamificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.IdUser == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Id User is required"})
		return
	}

	idUser, err := primitive.ObjectIDFromHex(req.IdUser.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Id User"})
		return
	}

	dataUser, err := repositories.GetGamificationByIdUser(context.Background(), idUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	if dataUser.Id != primitive.NilObjectID {
		c.JSON(http.StatusNotFound, dto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "User Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	gamification := dtoGamification.CreateGamificationRequest{
		Id:         primitive.NewObjectID(),
		IdUser:     idUser,
		Point:      req.Point,
		Challenges: []dtoGamification.Challenges{},
		DeletedAt:  nil,
		CreatedAt:  currentTime,
		CreatedBy:  "System",
		UpdatedAt:  currentTime,
		UpdatedBy:  "System",
	}

	result, err := repositories.CreateGamificationRequest(context.Background(), gamification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    result})
}

func GetAllGamificationHandler(c *gin.Context) {
	idUser := c.Query("_id_user")
	if idUser != "" {
		objectId, err := primitive.ObjectIDFromHex(idUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}
		members, err := repositories.GetGamificationByIdUser(context.Background(), objectId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.SuccessResult{
			Code:    http.StatusOK,
			Message: "success",
			Data:    members})
		return
	}

	gamifications, err := repositories.GetAllgamification(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    gamifications})
}

func UpdateGamificationHandler(c *gin.Context) {
	id := c.Query("_id")
	idUser := c.Query("_id_user")

	var objectId primitive.ObjectID
	var objectIdUser primitive.ObjectID
	var err error
	if id != "" {
		objectId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}

		dataById, err := repositories.GetGamificationById(context.Background(), objectId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}

		if dataById.Id == primitive.NilObjectID {
			c.JSON(http.StatusNotFound, dto.ErrorResult{
				Code:    http.StatusNotFound,
				Message: "Gamification not found"})
			return
		}
	}

	if idUser != "" {
		objectIdUser, err = primitive.ObjectIDFromHex(idUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}

		dataByIdUser, err := repositories.GetGamificationByIdUser(context.Background(), objectIdUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}

		if dataByIdUser.Id == primitive.NilObjectID {
			c.JSON(http.StatusNotFound, dto.ErrorResult{
				Code:    http.StatusNotFound,
				Message: "Gamification not found"})
			return
		}
	}

	var req dtoGamification.UpdatePointGamificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	updated := dtoGamification.UpdatePointGamificationRequest{
		Point:     req.Point,
		UpdatedAt: currentTime,
		UpdatedBy: "Admin",
	}

	updatedData, err := repositories.UpdateGamification(context.Background(), objectId, objectIdUser, updated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedData})
}

func ActiveGamificationHandler(c *gin.Context) {
	id := c.Query("_id")
	idUser := c.Query("_id_user")

	var objectId primitive.ObjectID
	var objectIdUser primitive.ObjectID
	var err error
	if id != "" {
		objectId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}

		dataById, err := repositories.GetGamificationById(context.Background(), objectId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}

		if dataById.Id == primitive.NilObjectID {
			c.JSON(http.StatusNotFound, dto.ErrorResult{
				Code:    http.StatusNotFound,
				Message: "Gamification not found"})
			return
		}
	}

	if idUser != "" {
		objectIdUser, err = primitive.ObjectIDFromHex(idUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}

		dataByIdUser, err := repositories.GetGamificationByIdUser(context.Background(), objectIdUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}

		if dataByIdUser.Id == primitive.NilObjectID {
			c.JSON(http.StatusNotFound, dto.ErrorResult{
				Code:    http.StatusNotFound,
				Message: "Gamification not found"})
			return
		}
	}

	var req dtoGamification.ActiveGamificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	updated := dtoGamification.ActiveGamificationRequest{
		DeletedAt: nil,
		UpdatedAt: currentTime,
		UpdatedBy: "Admin",
	}

	updatedData, err := repositories.ActiveGamification(context.Background(), objectId, objectIdUser, updated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedData})
}

func DeleteGamificationHandler(c *gin.Context) {
	id := c.Query("_id")
	idUser := c.Query("_id_user")

	var objectId primitive.ObjectID
	var objectIdUser primitive.ObjectID
	var err error
	if id != "" {
		objectId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}

		dataById, err := repositories.GetGamificationById(context.Background(), objectId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}

		if dataById.Id == primitive.NilObjectID {
			c.JSON(http.StatusNotFound, dto.ErrorResult{
				Code:    http.StatusNotFound,
				Message: "Gamification not found"})
			return
		}
	}

	if idUser != "" {
		objectIdUser, err = primitive.ObjectIDFromHex(idUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}

		dataByIdUser, err := repositories.GetGamificationByIdUser(context.Background(), objectIdUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}

		if dataByIdUser.Id == primitive.NilObjectID {
			c.JSON(http.StatusNotFound, dto.ErrorResult{
				Code:    http.StatusNotFound,
				Message: "Gamification not found"})
			return
		}
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	updated := dtoGamification.ActiveGamificationRequest{
		DeletedAt: &currentTime,
		UpdatedAt: currentTime,
		UpdatedBy: "Admin",
	}

	updatedData, err := repositories.ActiveGamification(context.Background(), objectId, objectIdUser, updated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedData})
}

func AddChallengesHandler(c *gin.Context) {
	id := c.Query("_id")
	idUser := c.Query("_id_user")

	if condition := id == "" && idUser == ""; condition {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "ID or ID User is required"})
		return
	}

	var objectId primitive.ObjectID
	var objectIdUser primitive.ObjectID
	var err error
	if id != "" {
		objectId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}

		dataById, err := repositories.GetGamificationById(context.Background(), objectId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}

		if dataById.Id == primitive.NilObjectID {
			c.JSON(http.StatusNotFound, dto.ErrorResult{
				Code:    http.StatusNotFound,
				Message: "Gamification not found"})
			return
		}
	}

	if idUser != "" {
		objectIdUser, err = primitive.ObjectIDFromHex(idUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}

		dataByIdUser, err := repositories.GetGamificationByIdUser(context.Background(), objectIdUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}

		if dataByIdUser.Id == primitive.NilObjectID {
			c.JSON(http.StatusNotFound, dto.ErrorResult{
				Code:    http.StatusNotFound,
				Message: "Gamification not found"})
			return
		}
	}

	var req dtoGamification.Challenges
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))
	// extra := "T00:00:00Z"

	challenges := dtoGamification.Challenges{
		Id:          primitive.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
		Point:       req.Point,
		Progress:    req.Progress,
		OnProgress:  req.OnProgress,
		Sponsor:     req.Sponsor,
		Claim:       false,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		DeletedAt:   nil,
		CreatedAt:   currentTime,
		CreatedBy:   "Admin",
		UpdatedAt:   currentTime,
		UpdatedBy:   "Admin",
	}

	updatedData, err := repositories.AddChallenge(context.Background(), objectId, objectIdUser, challenges)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedData})
}

func GetAllChallenges(c *gin.Context) {
	id := c.Query("_id")
	idGamification := c.Query("_id_gamification")
	idUser := c.Query("_id_user")

	if condition := id == "" && idUser == "" && idGamification == ""; condition {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "ID or ID User is required"})
		return
	}

	var objectId primitive.ObjectID
	var objectIdGamification primitive.ObjectID
	var objectIdUser primitive.ObjectID
	var err error

	if id != "" {
		objectId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}
		_, err := repositories.GetChallengesById(context.Background(), objectId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}
	}

	if idGamification != "" {
		objectIdGamification, err = primitive.ObjectIDFromHex(idGamification)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}
		_, err := repositories.GetGamificationById(context.Background(), objectIdGamification)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}
	}

	if idUser != "" {
		objectIdUser, err = primitive.ObjectIDFromHex(idUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}
		_, err := repositories.GetGamificationByIdUser(context.Background(), objectIdUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}
	}

	var Challenges []dtoGamification.Challenges
	var ChallengesObj dtoGamification.Challenges

	if condition := idGamification != "" || idUser != ""; condition {
		Challenges, err = repositories.GetChallenges(context.Background(), objectIdGamification, objectIdUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.SuccessResult{
			Code:    http.StatusOK,
			Message: "success",
			Data:    Challenges})
		return
	}

	if condition := id != ""; condition {
		ChallengesObj, err = repositories.GetChallengesById(context.Background(), objectId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.SuccessResult{
			Code:    http.StatusOK,
			Message: "success",
			Data:    ChallengesObj})
	}
}

func UpdateChallengesByIdHandler(c *gin.Context) {
	id := c.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format"})
		return
	}

	dataById, err := repositories.GetChallengesById(context.Background(), objectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	if dataById.Id == primitive.NilObjectID {
		c.JSON(http.StatusNotFound, dto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Challenge not found"})
		return
	}

	var req dtoGamification.Challenges
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	updated := dtoGamification.Challenges{}
	if req.Name != "" {
		updated.Name = req.Name
	}
	if req.Description != "" {
		updated.Description = req.Description
	}
	if req.Point != 0 {
		updated.Point = req.Point
	}
	if req.Progress != 0 {
		updated.Progress = req.Progress
	}
	if req.OnProgress != 0 {
		updated.OnProgress = req.OnProgress
	}
	if req.Sponsor != "" {
		updated.Sponsor = req.Sponsor
	}
	if req.Claim {
		updated.Claim = true
	}
	if !req.StartDate.IsZero() {
		updated.StartDate = req.StartDate
	}
	if !req.EndDate.IsZero() {
		updated.EndDate = req.EndDate
	}
	currentTime := time.Now()
	updated.UpdatedAt = currentTime
	updated.UpdatedBy = "Admin"

	challenges, err := repositories.UpdateChallengesById(context.Background(), objectId, updated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    challenges})
}
