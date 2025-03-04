package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type City struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	ProvinceID primitive.ObjectID `bson:"_id_province"`
	CreatedAt  time.Time          `bson:"created_at"`
	CreatedBy  string             `bson:"created_by"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	UpdatedBy  string             `bson:"updated_by"`
	DeletedAt  *time.Time         `bson:"deleted_at,omitempty"`
}
