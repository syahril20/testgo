package handlers

import (
	"context"
	"net/http"
	dtoHpi "server/dto/hpi"
	dto "server/dto/result"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllActiveHpiHandler(c *gin.Context) {
	hpis, err := repositories.GetAllHPIs(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    hpis})
}

func GetHpiHandlerById(c *gin.Context) {
	id := c.Param("id")
	hpiId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Hpi ID"})
		return
	}
	hpis, err := repositories.GetHPIById(context.Background(), hpiId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    hpis})
}

func GetAllNonActiveHpiHandler(c *gin.Context) {
	hpis, err := repositories.GetAllNonActiveHPIs(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    hpis})
}

func CreateHpiHandler(c *gin.Context) {
	var req dtoHpi.CreateHpiRequest
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

	hpiName, _ := repositories.GetHPIByName(context.Background(), req.Name)
	if hpiName != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	hpi := dtoHpi.CreateHpiRequest{
		Id:        primitive.NewObjectID(),
		Name:      req.Name,
		Biomarker: []dtoHpi.CreateBiomarkerRequest{},
		DeletedAt: nil,
		CreatedAt: currentTime,
		CreatedBy: "System",
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	data, err := repositories.CreateHPI(context.Background(), hpi)
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

func UpdateHpiHandler(c *gin.Context) {
	id := c.Param("id")
	var req dtoHpi.CreateHpiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Name == "" || id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	hpiId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Hpi ID"})
		return
	}

	existingHpi, _ := repositories.GetHPIById(context.Background(), hpiId)
	if existingHpi == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	hpiName, _ := repositories.GetHPIByName(context.Background(), req.Name)
	if hpiName != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	hpi := dtoHpi.UpdateHpiRequest{
		Name:      req.Name,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.UpdateHPIById(context.Background(), hpiId, hpi)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    hpi})
}

func DeleteHpiHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	hpiId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Hpi ID"})
		return
	}

	existingHpi, _ := repositories.GetHPIById(context.Background(), hpiId)
	if existingHpi == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	hpi := dtoHpi.ActiveHpiRequest{
		DeletedAt: &currentTime,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.ActiveHPIById(context.Background(), hpiId, hpi)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    hpi})
}

func ActiveHpiHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	hpiId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Hpi ID"})
		return
	}

	existingHpi, _ := repositories.GetHPIById(context.Background(), hpiId)
	if existingHpi == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	activeHpi := dtoHpi.ActiveHpiRequest{
		DeletedAt: nil,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.ActiveHPIById(context.Background(), hpiId, activeHpi)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    activeHpi})
}

// Biomarker
func GetAllActiveBiomarkersHandler(c *gin.Context) {
	id := c.Param("id")
	idHpi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Biomarker ID"})
		return
	}
	biomarkers, err := repositories.GetAllBiomarkers(context.Background(), idHpi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    biomarkers})
}

func GetBiomarkerHandlerById(c *gin.Context) {
	id := c.Query("id")
	biomarkerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Biomarker ID"})
		return
	}
	biomarker, err := repositories.GetBiomarkerById(context.Background(), biomarkerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    biomarker})
}

func GetAllNonActiveBiomarkersHandler(c *gin.Context) {
	id := c.Param("id")
	idHpi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Biomarker ID"})
		return
	}
	biomarkers, err := repositories.GetAllNonActiveBiomarkers(context.Background(), idHpi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    biomarkers})
}

func CreateBiomarkerHandler(c *gin.Context) {
	var req dtoHpi.CreateBiomarkerRequest
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

	hpiId, err := primitive.ObjectIDFromHex(req.IdHpi.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Biomarker ID"})
		return
	}

	biomarkerName, _ := repositories.GetBiomarkerByName(context.Background(), req.Name)
	if biomarkerName != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	biomarker := dtoHpi.CreateBiomarkerRequest{
		Id:        primitive.NewObjectID(),
		IdHpi:     req.IdHpi,
		Name:      req.Name,
		CreatedAt: currentTime,
		CreatedBy: "System",
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	data, err := repositories.CreateBiomarker(context.Background(), hpiId, biomarker)
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

func UpdateBiomarkerHandler(c *gin.Context) {
	id := c.Param("id")
	var req dtoHpi.CreateBiomarkerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Name == "" || id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	biomarkerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Biomarker ID"})
		return
	}

	existingBiomarker, _ := repositories.GetBiomarkerById(context.Background(), biomarkerId)
	if existingBiomarker == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	biomarkerName, _ := repositories.GetBiomarkerByName(context.Background(), req.Name)
	if biomarkerName != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	biomarker := dtoHpi.UpdateBiomarkerRequest{
		Name:      req.Name,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.UpdateBiomarkerById(context.Background(), biomarkerId, biomarker)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    biomarker})
}

func DeleteBiomarkerHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	biomarkerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Biomarker ID"})
		return
	}

	existingBiomarker, _ := repositories.GetBiomarkerById(context.Background(), biomarkerId)
	if existingBiomarker == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	biomarker := dtoHpi.ActiveBiomarkerRequest{
		DeletedAt: &currentTime,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.ActiveBiomarkerById(context.Background(), biomarkerId, biomarker)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    biomarker})
}

func ActiveBiomarkerHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	biomarkerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Biomarker ID"})
		return
	}

	existingBiomarker, _ := repositories.GetBiomarkerById(context.Background(), biomarkerId)
	if existingBiomarker == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	activeBiomarker := dtoHpi.ActiveBiomarkerRequest{
		DeletedAt: nil,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.ActiveBiomarkerById(context.Background(), biomarkerId, activeBiomarker)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    activeBiomarker})
}

// Under
func GetUnderByBiomarkerIdHandler(c *gin.Context) {
	id := c.Param("id")
	biomarkerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Biomarker ID"})
		return
	}

	under, err := repositories.GetUnderByBiomarkerId(context.Background(), biomarkerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    under})
}

func CreateUnderHandler(c *gin.Context) {
	id := c.Param("id")
	var req dtoHpi.CreateUnderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	biomarkerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Biomarker ID"})
		return
	}

	existingBiomarker, _ := repositories.GetBiomarkerById(context.Background(), biomarkerId)
	if existingBiomarker == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	existingUnder, _ := repositories.GetUnderByBiomarkerId(context.Background(), biomarkerId)
	if existingUnder != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	under := dtoHpi.CreateUnderRequest{
		Value:     req.Value,
		Unit:      req.Unit,
		Excercise: req.Excercise,
		Nutrision: req.Nutrision,
		CreatedAt: currentTime,
		CreatedBy: "System",
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	data, err := repositories.CreateUnder(context.Background(), biomarkerId, under)
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

func UpdateUnderHandler(c *gin.Context) {
	id := c.Param("id")
	var req dtoHpi.UpdateUnderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	biomarkerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Under ID"})
		return
	}

	existingUnder, _ := repositories.GetUnderByBiomarkerId(context.Background(), biomarkerId)
	if existingUnder == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	under := dtoHpi.UpdateUnderRequest{
		Value:     req.Value,
		Unit:      req.Unit,
		Excercise: req.Excercise,
		Nutrision: req.Nutrision,
		UpdatedAt: currentTime,
		UpdatedBy: "Admin",
	}

	_, err = repositories.UpdateUnderByBiomarkerId(context.Background(), biomarkerId, under)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    under})
}

func GetOverByBiomarkerIdHandler(c *gin.Context) {
	id := c.Param("id")
	biomarkerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Biomarker ID"})
		return
	}

	over, err := repositories.GetOverByBiomarkerId(context.Background(), biomarkerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    over})
}

func CreateOverHandler(c *gin.Context) {
	id := c.Param("id")
	var req dtoHpi.CreateOverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	biomarkerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Biomarker ID"})
		return
	}

	existingBiomarker, _ := repositories.GetBiomarkerById(context.Background(), biomarkerId)
	if existingBiomarker == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	existingOver, _ := repositories.GetOverByBiomarkerId(context.Background(), biomarkerId)
	if existingOver != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	over := dtoHpi.CreateOverRequest{
		Value:     req.Value,
		Unit:      req.Unit,
		Excercise: req.Excercise,
		Nutrision: req.Nutrision,
		CreatedAt: currentTime,
		CreatedBy: "System",
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	data, err := repositories.CreateOver(context.Background(), biomarkerId, over)
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

func UpdateOverHandler(c *gin.Context) {
	id := c.Param("id")
	var req dtoHpi.UpdateOverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	biomarkerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Over ID"})
		return
	}

	existingOver, _ := repositories.GetOverByBiomarkerId(context.Background(), biomarkerId)
	if existingOver == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	over := dtoHpi.UpdateOverRequest{
		Value:     req.Value,
		Unit:      req.Unit,
		Excercise: req.Excercise,
		Nutrision: req.Nutrision,
		UpdatedAt: currentTime,
		UpdatedBy: "Admin",
	}

	_, err = repositories.UpdateOverByBiomarkerId(context.Background(), biomarkerId, over)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    over})
}

func UpdateLifestyleHandler(c *gin.Context) {
	id := c.Param("id")
	var req dtoHpi.UpdateLifestyleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	biomarkerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Biomarker ID"})
		return
	}

	existingBiomarker, _ := repositories.GetBiomarkerById(context.Background(), biomarkerId)
	if existingBiomarker == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	biomarker := dtoHpi.UpdateLifestyleRequest{
		Lifestyle: req.Lifestyle,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.UpdateLifeStyleById(context.Background(), biomarkerId, biomarker)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    biomarker})
}

func CreateHpiResultsHandler(c *gin.Context) {
	var req dtoHpi.HpiPayload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.IdBiomarker == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Id Biomarker is required"})
		return
	}

	idBiomarker, err := primitive.ObjectIDFromHex(req.IdBiomarker.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Id Biomarker"})
		return
	}

	biomarkerData, err := repositories.GetBiomarkerById(context.Background(), idBiomarker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	if biomarkerData == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Biomarker data not found"})
		return
	}

	if req.Value == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Value is Required"})
		return
	}

	level := "Normal"
	excercise := ""
	nutrision := ""

	if biomarkerData.Under != nil && biomarkerData.Under.Value != "" && req.Value <= biomarkerData.Under.Value {
		level = "Under"
		excercise = biomarkerData.Under.Excercise
		nutrision = biomarkerData.Under.Nutrision
	} else if biomarkerData.Over != nil && biomarkerData.Over.Value != "" && req.Value >= biomarkerData.Over.Value {
		level = "Over"
		excercise = biomarkerData.Over.Excercise
		nutrision = biomarkerData.Over.Nutrision
	}

	if level == "Normal" {
		biomarkerData.Lifestyle = ""
	}

	currentTime := time.Now()

	result := dtoHpi.HpiResult{
		Id:            primitive.NewObjectID(),
		IdBiomarker:   idBiomarker,
		IdUser:        req.IdUser,
		NameBiomarker: biomarkerData.Name,
		Value:         req.Value,
		Level:         level,
		Excercise:     excercise,
		Nutrision:     nutrision,
		Lifestyle:     biomarkerData.Lifestyle,
		CheckDate:     req.CheckDate,
		CreatedAt:     currentTime,
		CreatedBy:     "Admin",
		UpdatedAt:     currentTime,
		UpdatedBy:     "Admin",
	}

	insertData, err := repositories.CreateHpiResult(context.Background(), result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    insertData})
}

func GetHpiResultHandlerById(c *gin.Context) {
	idUser := c.Param("id")
	ObjectIdUser, err := primitive.ObjectIDFromHex(idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Hpi ID"})
		return
	}
	hpiResult, err := repositories.GetHpiResultByIdUser(context.Background(), ObjectIdUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    hpiResult})
}
