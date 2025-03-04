package dto

import "time"

type CategoryRequest struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Image     string     `json:"image"`
	Icon      string     `json:"icon"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	CreatedBy string     `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
