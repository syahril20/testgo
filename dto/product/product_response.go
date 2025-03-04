package productDto

import "time"

type ProductResponse struct {
	Name      string    `json:"name" bson:"name"`
	Content   string    `json:"content" bson:"content"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at,omitempty"`
	Image     string    `json:"image" bson:"image"`
}
