package member

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateMemberRequest struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Price     int                `json:"price" bson:"price"`
	Benefit   string             `json:"benefit" bson:"benefit"`
	DeletedAt *time.Time         `json:"deleted_at" bson:"deleted_at"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by"`
}

type UpdateMemberRequest struct {
	Name      string    `json:"name" bson:"name"`
	Price     int       `json:"price" bson:"price"`
	Benefit   string    `json:"benefit" bson:"benefit"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type ActiveMemberRequest struct {
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
}
