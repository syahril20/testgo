package dto

import "time"

type CreateFitnessProgramRequest struct {
	ID          string     `json:"id,omitempty"`
	Title       string     `json:"title" binding:"required"`
	Image       string     `json:"image" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Instructor  string     `json:"instructor" binding:"required"`
	Date        string     `json:"date" binding:"required"`
	Trait       []string   `json:"trait"`
	Duration    string     `json:"duration" binding:"required"`
	Category    string     `json:"category" binding:"required"`
	Link        string     `json:"link" binding:"required"`
	CreatedAt   time.Time  `json:"created_at" bson:"created_at"`
	CreatedBy   string     `json:"created_by" bson:"created_by"`
	DeletedAt   *time.Time `bson:"deleted_at,omitempty"`
}

type UpdateFitnessProgramRequest struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Image       string    `json:"image,omitempty"`
	Description string    `json:"description,omitempty"`
	Instructor  string    `json:"instructor,omitempty"`
	Date        string    `json:"date,omitempty"`
	Trait       []string  `json:"trait,omitempty"`
	Duration    string    `json:"duration,omitempty"`
	Category    string    `json:"category,omitempty"`
	Link        string    `json:"link,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy   string    `json:"updated_by" bson:"updated_by"`
}
