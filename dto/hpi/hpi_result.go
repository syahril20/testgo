package hpi

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HpiPayload struct {
	IdBiomarker primitive.ObjectID `json:"_id_biomarker" bson:"_id_biomarker"`
	IdUser      primitive.ObjectID `json:"_id_user" bson:"_id_user"`
	Value       string             `json:"value" bson:"value"`
	CheckDate   time.Time          `json:"check_date" bson:"check_date"`
}

type HpiResult struct {
	Id            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	IdHpi         primitive.ObjectID `json:"_id_hpi" bson:"_id_hpi"`
	IdBiomarker   primitive.ObjectID `json:"_id_biomarker" bson:"_id_biomarker"`
	IdUser        primitive.ObjectID `json:"_id_user" bson:"_id_user"`
	NameHpi       string             `json:"name_hpi" bson:"name_hpi"`
	NameBiomarker string             `json:"name_biomarker" bson:"name_biomarker"`
	Value         string             `json:"value" bson:"value"`
	Level         string             `json:"level" bson:"level"`
	Excercise     string             `json:"excercise" bson:"excercise"`
	Nutrision     string             `json:"nutrision" bson:"nutrision"`
	Lifestyle     string             `json:"lifestyle" bson:"lifestyle"`
	CheckDate     time.Time          `json:"check_date" bson:"check_date"`
	DeletedAt     *time.Time         `json:"deleted_at" bson:"deleted_at"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy     string             `json:"created_by" bson:"created_by"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy     string             `json:"updated_by" bson:"updated_by"`
}
