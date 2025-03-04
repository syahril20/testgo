package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // ID kategori
	Title     string             `bson:"title"`         // Nama kategori
	Image     string             `bson:"image"`         // URL atau path gambar
	Icon      string             `bson:"icon"`          // URL atau path ikon
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by"`
}
