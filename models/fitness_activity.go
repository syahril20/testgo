package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FitnessActivity struct {
	ID        primitive.ObjectID `bson:"_id"`
	IDuser    primitive.ObjectID `bson:"_id_user"`
	IDfitness primitive.ObjectID `bson:"_id_fitness"`
	Finished  bool               `bson:"finished"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by"`
}
