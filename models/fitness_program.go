package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FitnessProgram struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Image       string             `bson:"image" json:"image"`
	Description string             `bson:"description" json:"description"`
	Instructor  string             `bson:"instructor" json:"instructor"`
	Date        time.Time          `bson:"date" json:"date"`
	Trait       []string           `bson:"trait" json:"trait"`
	Duration    string             `bson:"duration" json:"duration"`
	Category    string             `bson:"category" json:"category"`
	Link        string             `bson:"link" json:"link"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy   string             `json:"updated_by" bson:"updated_by"`
	DeletedAt   *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
