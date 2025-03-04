package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PaymentID primitive.ObjectID `bson:"payment_id" json:"payment_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	CreatedBy string             `bson:"created_by" json:"created_by"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	UpdatedBy string             `bson:"updated_by" json:"updated_by"`
}
