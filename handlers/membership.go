package handlers

import (
	"context"
	"net/http"
	dtoMembership "server/dto/membership"
	dto "server/dto/result"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMembershipByIdHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "ID is required"})
		return
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format"})
		return
	}

	membership, err := repositories.GetMembershipByID(context.Background(), objectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    membership})
}

func GetAllMembershipHandler(c *gin.Context) {
	id := c.Query("_id")
	idMember := c.Query("_id_member")
	idUser := c.Query("_id_user")

	if id != "" {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}
		members, err := repositories.GetMembershipByID(context.Background(), objectId)
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

	} else if idMember != "" {
		objectId, err := primitive.ObjectIDFromHex(idMember)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}
		members, err := repositories.GetMembershipByIdMember(context.Background(), objectId)
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

	} else if idUser != "" {
		objectId, err := primitive.ObjectIDFromHex(idUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}
		members, err := repositories.GetMembershipByIdUser(context.Background(), objectId)
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

	members, err := repositories.GetAllMembership(context.Background())
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
}

func CreateMembershipHandler(c *gin.Context) {
	var req dtoMembership.CreateMembershipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.IdMember.Hex() == "" || req.IdUser.Hex() == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	_, err := primitive.ObjectIDFromHex(req.IdMember.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid IdMember ID"})
		return
	}

	_, err = primitive.ObjectIDFromHex(req.IdUser.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid IdUser ID"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))
	oneMonth := currentTime.AddDate(0, 1, 0)

	membership := dtoMembership.CreateMembershipRequest{
		Id:        primitive.NewObjectID(),
		IdMember:  req.IdMember,
		IdUser:    req.IdUser,
		EndDate:   oneMonth,
		DeletedAt: nil,
		CreatedAt: currentTime,
		CreatedBy: "System",
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	data, err := repositories.CreateMembership(context.Background(), membership)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    data})
}

func UpdateMembershipByIdHandler(c *gin.Context) {
	var req dtoMembership.UpdateMembershipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "ID is required"})
		return
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format"})
		return
	}

	membership, err := repositories.GetMembershipByID(context.Background(), objectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	if membership.Id.IsZero() {
		c.JSON(http.StatusNotFound, dto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Member not found"})
		return
	}

	var membershipUpdate dtoMembership.UpdateMembershipRequest

	if req.IdMember != "" {
		_, err := primitive.ObjectIDFromHex(req.IdMember)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}
		membershipUpdate.IdMember = req.IdMember
	}
	if req.IdUser != "" {
		_, err := primitive.ObjectIDFromHex(req.IdUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Invalid ID format"})
			return
		}
		membershipUpdate.IdUser = req.IdUser
	}

	membershipUpdate.UpdatedAt = time.Now().In(time.FixedZone("UTC+7", 7*3600))
	membershipUpdate.UpdatedBy = "Admin"

	updatedMembership, err := repositories.UpdateMembershipByID(context.Background(), objectId, membershipUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedMembership})
}

func ActiveMembershipByIdHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "ID is required"})
		return
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format"})
		return
	}

	membership, err := repositories.GetMembershipByID(context.Background(), objectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	if membership.Id.IsZero() {
		c.JSON(http.StatusNotFound, dto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Member not found"})
		return
	}

	updateData := dtoMembership.ActiveMembershipRequest{}

	curentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	updateData.DeletedAt = nil
	updateData.UpdatedAt = curentTime
	updateData.UpdatedBy = "Admin"

	updatedMembership, err := repositories.ActiveMembershipByID(context.Background(), objectId, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedMembership})
}

func DeleteMembershipHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "ID is required"})
		return
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format"})
		return
	}

	membership, err := repositories.GetMembershipByID(context.Background(), objectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	if membership.Id.IsZero() {
		c.JSON(http.StatusNotFound, dto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Member not found"})
		return
	}

	updateData := dtoMembership.ActiveMembershipRequest{}

	curentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	updateData.DeletedAt = &curentTime
	updateData.UpdatedAt = curentTime
	updateData.UpdatedBy = "Admin"

	updatedMembership, err := repositories.ActiveMembershipByID(context.Background(), objectId, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedMembership})
}
