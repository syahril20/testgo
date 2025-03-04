package handlers

import (
	"context"
	"net/http"
	dtoassessment "server/dto/assessment"
	dto "server/dto/result"
	"server/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// assessment

func GetAllActiveassessmentsHandler(c *gin.Context) {
	assessments, err := repositories.GetAllassessments(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    assessments})
}

func GetAllNonActiveassessmentsHandler(c *gin.Context) {
	assessments, err := repositories.GetAllNonActiveassessments(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    assessments})
}

func CreateassessmentHandler(c *gin.Context) {
	var req dtoassessment.CreateassessmentRequest
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
			Message: "Invalid Product ID"})
		return
	}

	assessmentName, _ := repositories.GetassessmentByName(context.Background(), req.Name)
	if assessmentName != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	assessment := dtoassessment.CreateassessmentRequest{
		Id:            primitive.NewObjectID(),
		Name:          req.Name,
		Code:          req.Code,
		DeletedAt:     nil,
		Questionnaire: []dtoassessment.CreateQuestionnaireRequest{},
		CreatedAt:     currentTime,
		CreatedBy:     "System",
		UpdatedAt:     currentTime,
		UpdatedBy:     "System",
	}

	data, err := repositories.Createassessment(context.Background(), assessment)
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

func UpdateassessmentHandler(c *gin.Context) {
	id := c.Param("id")
	var req dtoassessment.CreateassessmentRequest
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

	assessmentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Product ID"})
		return
	}

	existingassessment, _ := repositories.GetassessmentById(context.Background(), assessmentId)
	if existingassessment == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	assessmentName, _ := repositories.GetassessmentByName(context.Background(), req.Name)
	if assessmentName != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	assessment := dtoassessment.UpdateassessmentRequest{
		Name:      req.Name,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.UpdateassessmentById(context.Background(), assessmentId, assessment)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    assessment})
}

func DeleteassessmentHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	questionnaireId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid assessment ID"})
		return
	}

	existingassessment, _ := repositories.GetassessmentById(context.Background(), questionnaireId)
	if existingassessment == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	assessment := dtoassessment.ActiveassessmentRequest{
		DeletedAt: &currentTime,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.ActiveassessmentById(context.Background(), questionnaireId, assessment)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    assessment})
}

func ActiveassessmentHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	assessmentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid assessment ID"})
		return
	}

	existingassessment, _ := repositories.GetassessmentById(context.Background(), assessmentId)
	if existingassessment == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	activeassessment := dtoassessment.ActiveassessmentRequest{
		DeletedAt: nil,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.ActiveassessmentById(context.Background(), assessmentId, activeassessment)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    activeassessment})
}

// Questionnaire

func GetAllActiveQuestionnaireHandler(c *gin.Context) {
	id := c.Param("id")
	assessmentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid assessment ID"})
		return
	}
	assessments, err := repositories.GetAllDataQuestionnaires(context.Background(), assessmentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    assessments})
}

func GetAllNonActiveQuestionnaireHandler(c *gin.Context) {
	id := c.Param("id")
	assessmentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid assessment ID"})
		return
	}
	assessments, err := repositories.GetAllDataQuestionnaires(context.Background(), assessmentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    assessments})
}

func CreateQuestionnaireHandler(c *gin.Context) {
	var req dtoassessment.CreateQuestionnaireRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Question == "" || req.IdAsessment.Hex() == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	assessmentId, err := primitive.ObjectIDFromHex(req.IdAsessment.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Product ID"})
		return
	}

	existingQuestionnaire, _ := repositories.GetassessmentById(context.Background(), assessmentId)
	if existingQuestionnaire == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data assessment Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	newQuestionnaire := dtoassessment.CreateQuestionnaireRequest{
		Id:          primitive.NewObjectID(),
		IdAsessment: req.IdAsessment,
		Question:    req.Question,
		Code:        req.Code,
		CreatedAt:   currentTime,
		CreatedBy:   "System",
		UpdatedAt:   currentTime,
		UpdatedBy:   "System",
	}

	data, err := repositories.CreateQuestionnaire(context.Background(), assessmentId, newQuestionnaire)
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

func UpdateQuestionnaireHandler(c *gin.Context) {
	id := c.Param("id")
	var req dtoassessment.UpdateQuestionnaireRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Question == "" || id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	questionnaireId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Questionnaire ID"})
		return
	}

	existingQuestionnaire, _ := repositories.GetQuestionnaireById(context.Background(), questionnaireId)
	if existingQuestionnaire == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	nameQuestion, _ := repositories.GetQuestionnaireByQuestion(context.Background(), req.Question)
	if nameQuestion != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	updatedQuestionnaire := dtoassessment.UpdateQuestionnaireRequest{
		Question:  req.Question,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.UpdateQuestionnaireById(context.Background(), questionnaireId, updatedQuestionnaire)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedQuestionnaire})
}

func DeleteQuestionnaireHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	questionnaireId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Questionnaire ID"})
		return
	}

	existingQuestionnaire, _ := repositories.GetQuestionnaireById(context.Background(), questionnaireId)
	if existingQuestionnaire == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	softDeletedQuestionnaire := dtoassessment.ActiveQuestionnaireRequest{
		DeletedAt: &currentTime,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.ActiveQuestionnaireById(context.Background(), questionnaireId, softDeletedQuestionnaire)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    softDeletedQuestionnaire})
}

func ActiveQuestionnaireHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	questionnaireId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Questionnaire ID"})
		return
	}

	existingQuestionnaire, _ := repositories.GetQuestionnaireById(context.Background(), questionnaireId)
	if existingQuestionnaire == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	ActiveQuestionnaire := dtoassessment.ActiveQuestionnaireRequest{
		DeletedAt: nil,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.ActiveQuestionnaireById(context.Background(), questionnaireId, ActiveQuestionnaire)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    ActiveQuestionnaire})
}

// Assesment Payload

func GetAssessmentPayloadHandler(c *gin.Context) {
	var req dtoassessment.AssesementPayload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if len(req.QuestionnairePayload) == 0 {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "QuestionnairePayload cannot be empty"})
		return
	}

	var questionnairePayloads []dtoassessment.QuestionnairePayload
	for _, qp := range req.QuestionnairePayload {
		questionnairePayloads = append(questionnairePayloads, dtoassessment.QuestionnairePayload{
			IdQuestionnaire: qp.IdQuestionnaire,
			Code:            qp.Code,
			Answer:          qp.Answer,
		})
	}

	assessmentId, err := primitive.ObjectIDFromHex(req.IdAsessment.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid assessment ID"})
		return
	}

	existingAssessment, _ := repositories.GetassessmentById(context.Background(), assessmentId)
	if existingAssessment == nil {
		c.JSON(http.StatusNotFound, dto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "Data Not Found"})
		return
	}

	var result string

	if existingAssessment.Code == "SERVIKS" {
		result = CheckServiks(assessmentId, questionnairePayloads)
	} else if existingAssessment.Code == "BREAST" {
		result = CheckBreast(assessmentId, questionnairePayloads)
	} else if existingAssessment.Code == "LIVER" {
		result = CheckLiver(assessmentId, questionnairePayloads)
	} else if existingAssessment.Code == "LUNG" {
		result = CheckLung(assessmentId, questionnairePayloads)
	} else if existingAssessment.Code == "KOLOREKTAL" {
		result = CheckKolorektal(assessmentId, questionnairePayloads)
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    result})
}

func CheckServiks(assessmentId primitive.ObjectID, questionnairePayloads []dtoassessment.QuestionnairePayload) string {
	minAge := 21
	maxAge := 65
	trueCount := 0
	for _, qp := range questionnairePayloads {
		if qp.Code == "AGE" {
			age, err := strconv.Atoi(qp.Answer)
			if err != nil {
				return "Invalid age format"
			}
			if age >= minAge && age <= maxAge {
				trueCount++
			}
		}
	}

	for _, qp := range questionnairePayloads {
		if qp.Answer == "true" {
			trueCount++
		}
	}

	result := "low"
	if trueCount > 1 {
		result = "high"
	}
	return result
}

func CheckBreast(assessmentId primitive.ObjectID, questionnairePayloads []dtoassessment.QuestionnairePayload) string {
	maxAge := 30
	var overAge bool
	trueCount := 0
	result := "low"
	for _, qp := range questionnairePayloads {
		if qp.Code == "AGE" {
			age, err := strconv.Atoi(qp.Answer)
			if err != nil {
				return "Invalid age format"
			}
			if age <= maxAge {
				trueCount++
				overAge = false

			} else if age >= maxAge {
				trueCount++
				overAge = true
			}
		}
	}

	for _, qp := range questionnairePayloads {
		if qp.Answer == "true" {
			trueCount++
		}
	}

	if trueCount > 1 && !overAge {
		result = "high (MRI payudara)"
	} else if trueCount > 1 && overAge {
		result = "high (USG payudara / mammografi)"
	}

	return result
}

func CheckLiver(assessmentId primitive.ObjectID, questionnairePayloads []dtoassessment.QuestionnairePayload) string {
	trueCount := 0
	result := "low"
	for _, qp := range questionnairePayloads {
		if qp.Answer == "true" {
			trueCount++
		}
	}

	if trueCount > 1 {
		result = "high"
	}

	return result
}

func CheckLung(assessmentId primitive.ObjectID, questionnairePayloads []dtoassessment.QuestionnairePayload) string {
	trueCount := 0
	smoker := false
	under := false
	packs := 0
	years := 0
	overAge := false
	overAges := false
	result := "low"

	for _, qp := range questionnairePayloads {
		if qp.Code == "AGE" {
			age, err := strconv.Atoi(qp.Answer)
			if err != nil {
				return "Invalid age format"
			}
			if age > 40 {
				overAge = true
			}
			if age > 50 {
				overAges = true
			}
		}
		if qp.Code == "SMOKER" && qp.Answer == "true" {
			smoker = true
		}

		if qp.Code == "SMOKING" && qp.Answer == "true" {
			under = true
		} else if qp.Code == "SMOKING" && qp.Answer == "false" {
			under = false
		}

		if qp.Code == "PACK" {
			pack, err := strconv.Atoi(qp.Answer)
			if err != nil {
				return "Invalid format"
			}
			packs = pack
		}
		if qp.Code == "YEARS" {
			year, err := strconv.Atoi(qp.Answer)
			if err != nil {
				return "Invalid format"
			}
			years = year
		}
	}
	if (overAge) && (smoker && under) && (packs*years >= 30) {
		result = "high"
		return result
	}

	for _, qp := range questionnairePayloads {
		if qp.Code == "SECOND" {
			if qp.Answer == "true" {
				trueCount++
			}
		}
	}

	if (trueCount >= 1) && (packs*years >= 20) && (overAges) {
		result = "high"
		return result
	}

	for _, qp := range questionnairePayloads {
		if qp.Code == "THIRD" {
			if qp.Answer == "true" {
				result = "high"
				return result
			}
		}
	}
	return result
}

func CheckKolorektal(assessmentId primitive.ObjectID, questionnairePayloads []dtoassessment.QuestionnairePayload) string {
	trueCount := 0
	result := "low"
	for _, qp := range questionnairePayloads {
		if qp.Answer == "true" {
			trueCount++
		}
	}

	if trueCount >= 1 {
		result = "high"
	}

	return result
}
