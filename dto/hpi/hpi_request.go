package hpi

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateHpiRequest struct {
	Id        primitive.ObjectID       `json:"_id" bson:"_id,omitempty"`
	Name      string                   `json:"name" bson:"name"`
	Biomarker []CreateBiomarkerRequest `json:"biomarker" bson:"biomarker"`
	DeletedAt *time.Time               `json:"deleted_at" bson:"deleted_at"`
	CreatedAt time.Time                `json:"created_at" bson:"created_at"`
	CreatedBy string                   `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time                `json:"updated_at" bson:"updated_at"`
	UpdatedBy string                   `json:"updated_by" bson:"updated_by"`
}

type UpdateHpiRequest struct {
	Name      string    `json:"name" bson:"name"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type ActiveHpiRequest struct {
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
}

type CreateBiomarkerRequest struct {
	Id        primitive.ObjectID  `json:"_id" bson:"_id,omitempty"`
	IdHpi     primitive.ObjectID  `json:"_id_hpi" bson:"_id_hpi"`
	Name      string              `json:"name" bson:"name"`
	Under     *CreateUnderRequest `json:"under" bson:"under"`
	Over      *CreateOverRequest  `json:"over" bson:"over"`
	Lifestyle string              `json:"lifestyle" bson:"lifestyle"`
	DeletedAt *time.Time          `json:"deleted_at" bson:"deleted_at"`
	CreatedAt time.Time           `json:"created_at" bson:"created_at"`
	CreatedBy string              `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time           `json:"updated_at" bson:"updated_at"`
	UpdatedBy string              `json:"updated_by" bson:"updated_by"`
}

type UpdateBiomarkerRequest struct {
	Name      string    `json:"name" bson:"name"`
	Lifestyle string    `json:"lifestyle" bson:"lifestyle"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type UpdateLifestyleRequest struct {
	Lifestyle string    `json:"lifestyle" bson:"lifestyle"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type ActiveBiomarkerRequest struct {
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
}

type CreateUnderRequest struct {
	Value     string    `json:"value" bson:"value"`
	Unit      string    `json:"unit" bson:"unit"`
	Excercise string    `json:"excercise" bson:"excercise"`
	Nutrision string    `json:"nutrision" bson:"nutrision"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	CreatedBy string    `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type UpdateUnderRequest struct {
	Value     string    `json:"value" bson:"value"`
	Unit      string    `json:"unit" bson:"unit"`
	Excercise string    `json:"excercise" bson:"excercise"`
	Nutrision string    `json:"nutrision" bson:"nutrision"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type CreateOverRequest struct {
	Value     string    `json:"value" bson:"value"`
	Unit      string    `json:"unit" bson:"unit"`
	Excercise string    `json:"excercise" bson:"excercise"`
	Nutrision string    `json:"nutrision" bson:"nutrision"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	CreatedBy string    `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type UpdateOverRequest struct {
	Value     string    `json:"value" bson:"value"`
	Unit      string    `json:"unit" bson:"unit"`
	Excercise string    `json:"excercise" bson:"excercise"`
	Nutrision string    `json:"nutrision" bson:"nutrision"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}
