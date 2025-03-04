package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Province struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at"`
}
