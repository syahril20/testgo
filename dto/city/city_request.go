package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CityRequest struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ProvinceID string             `json:"_id_province" bson:"_id_province"`
	Name       string             `json:"name" bson:"name"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy  string             `json:"created_by" bson:"created_by"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy  string             `json:"updated_by" bson:"updated_by"`
	DeletedAt  *time.Time         `bson:"deleted_at,omitempty"`
}
