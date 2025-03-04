package dto

import (
	"time"
)

type CreateWellnessCategoryRequest struct {
	Content   string     `bson:"content"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
	Image     string     `bson:"image"`
	Title     string     `bson:"title"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	CreatedBy string     `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
}
