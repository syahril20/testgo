package gamification

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateGamificationRequest struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	IdUser     primitive.ObjectID `bson:"_id_user,omitempty" json:"_id_user"`
	Point      int                `bson:"point" json:"point"`
	Challenges []Challenges       `bson:"challenges" json:"challenges"`
	DeletedAt  *time.Time         `json:"deleted_at" bson:"deleted_at"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy  string             `json:"created_by" bson:"created_by"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy  string             `json:"updated_by" bson:"updated_by"`
}

type UpdatePointGamificationRequest struct {
	Point     int       `bson:"point" json:"point"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type ActiveGamificationRequest struct {
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
}

type Challenges struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Point       int                `bson:"point" json:"point"`
	Progress    int                `bson:"progress" json:"progress"`
	OnProgress  int                `bson:"on_progress" json:"on_progress"`
	Sponsor     string             `bson:"sponsor" json:"sponsor"`
	Claim       bool               `bson:"claim" json:"claim"`
	StartDate   time.Time          `bson:"start_date" json:"start_date"`
	EndDate     time.Time          `bson:"end_date" json:"end_date"`
	DeletedAt   *time.Time         `json:"deleted_at" bson:"deleted_at"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy   string             `json:"updated_by" bson:"updated_by"`
}
