package dto

import "time"

type AreaRequest struct {
	ID        string     `json:"_id" bson:"_id"`
	SuburbID  string     `json:"_id_suburb" bson:"_id_suburb"`
	Name      string     `json:"name" bson:"name"`
	Zip       string     `json:"zip" bson:"zip"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	CreatedBy string     `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
