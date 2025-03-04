package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"_id"`
	UserID    primitive.ObjectID  `bson:"_id_user" json:"_id_user"`
	ProductID primitive.ObjectID  `bson:"_id_product" json:"_id_product"`
	CreatedAt primitive.DateTime  `bson:"created_at" json:"created_at"`
	CreatedBy string              `bson:"created_by" json:"created_by"`
	UpdatedAt *primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	UpdatedBy *string             `bson:"updated_by,omitempty" json:"updated_by,omitempty"`
	DeletedAt *primitive.DateTime `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
