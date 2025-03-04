package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Area struct {
	ID        primitive.ObjectID  `bson:"id" json:"id"`
	SuburbID  primitive.ObjectID  `bson:"id_suburb" json:"suburb_id"`
	Name      string              `bson:"name" json:"name"`
	Zip       string              `bson:"zip" json:"zip"`
	CreatedAt primitive.DateTime  `bson:"created_at" json:"created_at"`
	CreatedBy string              `bson:"created_by" json:"created_by"`
	UpdatedAt primitive.DateTime  `bson:"updated_at" json:"updated_at"`
	UpdatedBy string              `bson:"updated_by" json:"updated_by"`
	DeletedAt *primitive.DateTime `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
