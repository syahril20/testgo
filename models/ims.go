package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IMS struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email                string             `bson:"email" json:"email"`
	Name                 string             `bson:"name" json:"name"`
	Old                  int                `bson:"old" json:"old"`
	Phone                int                `bson:"phone" json:"phone"`
	Address              string             `bson:"address" json:"address"`
	OptiSampleCollection string             `bson:"opti_sample_collection" json:"opti_sample_collection"`
	CreatedAt            time.Time          `bson:"created_at" json:"created_at"`
	CreatedBy            string             `bson:"created_by" json:"created_by"`
	UpdatedAt            time.Time          `bson:"updated_at" json:"updated_at"`
	UpdatedBy            string             `bson:"updated_by" json:"updated_by"`
	DeletedAt            time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}
