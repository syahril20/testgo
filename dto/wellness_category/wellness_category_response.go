package dto

import "time"

type WellnessCategoryResponse struct {
	Title     string     `bson:"title"`
	Content   string     `bson:"content"`
	Image     string     `bson:"image"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
