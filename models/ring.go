package models

import "time"

type Ring struct {
	ID         string    `json:"id" bson:"_id,omitempty"`
	Size       int       `json:"size" bson:"size"`
	Color      string    `json:"color" bson:"color"`
	Connection bool      `json:"connection" bson:"connection"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	CreatedBy  string    `json:"created_by" bson:"created_by"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy  string    `json:"updated_by" bson:"updated_by"`
}
