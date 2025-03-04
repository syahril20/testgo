package questionnaire

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Assesemnt
type CreateassessmentRequest struct {
	Id            primitive.ObjectID           `json:"_id" bson:"_id,omitempty"`
	Name          string                       `json:"name" bson:"name"`
	Code          string                       `json:"code" bson:"code"`
	DeletedAt     *time.Time                   `json:"deleted_at" bson:"deleted_at"`
	Questionnaire []CreateQuestionnaireRequest `json:"questionnaire" bson:"questionnaire"`
	CreatedAt     time.Time                    `json:"created_at" bson:"created_at"`
	CreatedBy     string                       `json:"created_by" bson:"created_by"`
	UpdatedAt     time.Time                    `json:"updated_at" bson:"updated_at"`
	UpdatedBy     string                       `json:"updated_by" bson:"updated_by"`
}

type UpdateassessmentRequest struct {
	Name      string    `json:"name" bson:"name"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type ActiveassessmentRequest struct {
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
}

// Questionnaire
type CreateQuestionnaireRequest struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	IdAsessment primitive.ObjectID `json:"_id_assessment" bson:"_id_assessment"`
	Code        string             `json:"code" bson:"code"`
	Question    string             `json:"question" bson:"question"`

	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	CreatedBy string     `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
}

type UpdateQuestionnaireRequest struct {
	Question  string    `json:"question" bson:"question"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type ActiveQuestionnaireRequest struct {
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
}

type AssesementPayload struct {
	IdAsessment          primitive.ObjectID     `json:"_id_assessment" bson:"_id_assessment"`
	QuestionnairePayload []QuestionnairePayload `json:"questionnaire" bson:"questionnaire"`
}

type QuestionnairePayload struct {
	IdQuestionnaire primitive.ObjectID `json:"_id_questionnaire" bson:"_id_questionnaire"`
	Code            string             `json:"code" bson:"code"`
	Answer          string             `json:"answer" bson:"answer"`
}
