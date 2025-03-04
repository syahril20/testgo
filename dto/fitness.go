package dto

import "time"

// import "go.mongodb.org/mongo-driver/bson/primitive"

type FitnessDTO struct {
	ID          string     `json:"id,omitempty" bson:"_id,omitempty"`
	CategoryId  string     `json:"categoryId"`
	Title       string     `json:"title"`
	Image       string     `json:"image"`
	Description string     `json:"description"`
	Video       []VideoDTO `json:"video"`
	Workout     string     `json:"workout"`
	CreatedAt   time.Time  `json:"created_at" bson:"created_at"`
	CreatedBy   string     `json:"created_by" bson:"created_by"`
	UpdatedAt   time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy   string     `json:"updated_by" bson:"updated_by"`
}

type VideoDTO struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	Duration  string `json:"duration"`
	Thumbnail string `json:"thumbnail"`
}
