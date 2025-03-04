package handlers

import (
	"context"
	"net/http"
	dtoMember "server/dto/member"
	dto "server/dto/result"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMembersHandler(c *gin.Context) {
	members, err := repositories.GetActiveMembers(context.Background())
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

func GetNonActiveAllMembersHandler(c *gin.Context) {
	members, err := repositories.GetNonActiveMembers(context.Background())
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

func GetMemberByIDHandler(c *gin.Context) {
	id := c.Param("id")
	memberID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Member ID"})
		return
	}

	member, err := repositories.GetMemberByID(context.Background(), memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	if member.Id.IsZero() {
		c.JSON(http.StatusNotFound, dto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Member not found"})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    member})
}

func CreateMemberHandler(c *gin.Context) {
	var req dtoMember.CreateMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	_, err := primitive.ObjectIDFromHex(req.Id.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Hpi ID"})
		return
	}

	memberName, _ := repositories.GetMemberByName(context.Background(), req.Name)
	if memberName != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	member := dtoMember.CreateMemberRequest{
		Id:        primitive.NewObjectID(),
		Name:      req.Name,
		Price:     req.Price,
		Benefit:   req.Benefit,
		DeletedAt: nil,
		CreatedAt: currentTime,
		CreatedBy: "System",
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	data, err := repositories.CreateMember(context.Background(), member)
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

func UpdateMemberByIDHandler(c *gin.Context) {
	id := c.Param("id")
	memberID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Member ID"})
		return
	}

	var req dtoMember.UpdateMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	member, err := repositories.GetMemberByID(context.Background(), memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	if member.Id.IsZero() {
		c.JSON(http.StatusNotFound, dto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Member not found"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	var updateReq dtoMember.UpdateMemberRequest

	if req.Name != "" {
		updateReq.Name = req.Name
	}
	if req.Price != 0 {
		updateReq.Price = req.Price
	}
	if req.Benefit != "" {
		updateReq.Benefit = req.Benefit
	}
	updateReq.UpdatedAt = currentTime
	updateReq.UpdatedBy = "Admin"

	updatedMember, err := repositories.UpdateMemberByID(context.Background(), memberID, updateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedMember})
}

func ActiveMemberHandler(c *gin.Context) {
	id := c.Param("id")
	memberID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Member ID"})
		return
	}

	member, err := repositories.GetMemberByID(context.Background(), memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	if member.Id.IsZero() {
		c.JSON(http.StatusNotFound, dto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Member not found"})
		return
	}

	updateData := dtoMember.ActiveMemberRequest{}

	updateData.DeletedAt = nil
	updateData.UpdatedAt = time.Now().In(time.FixedZone("UTC+7", 7*3600))
	updateData.UpdatedBy = "Admin"

	updatedMember, err := repositories.ActiveMemberByID(context.Background(), memberID, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedMember})
}

func DeleteMemberHandler(c *gin.Context) {
	id := c.Param("id")
	memberID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Member ID"})
		return
	}

	member, err := repositories.GetMemberByID(context.Background(), memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	if member.Id.IsZero() {
		c.JSON(http.StatusNotFound, dto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Member not found"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	updateData := dtoMember.ActiveMemberRequest{}

	updateData.DeletedAt = &currentTime
	updateData.UpdatedAt = currentTime
	updateData.UpdatedBy = "Admin"

	updatedMember, err := repositories.ActiveMemberByID(context.Background(), memberID, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedMember})
}
