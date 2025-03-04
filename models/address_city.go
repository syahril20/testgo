package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddressCity merepresentasikan data kota/kabupaten.
type AddressCity struct {
	ID         primitive.ObjectID `bson:"_id"`
	IDProvince primitive.ObjectID `bson:"_id_province"`
	Name       string             `bson:"name"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy  string             `json:"created_by" bson:"created_by"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy  string             `json:"updated_by" bson:"updated_by"`
}
