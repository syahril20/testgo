package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Content    string             `json:"content" bson:"content"`
	DeletedAt  bool               `json:"deleted_at" bson:"deleted_at,omitempty"`
	Image      string             `json:"image" bson:"image"`
	SubProduct []SubProduct       `json:"sub_product" bson:"sub_product"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy  string             `json:"created_by" bson:"created_by"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy  string             `json:"updated_by" bson:"updated_by"`
}

type SubProduct struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Content   string             `json:"content" bson:"content"`
	DeletedAt bool               `json:"deleted_at" bson:"deleted_at,omitempty"`
	Image     string             `json:"image" bson:"image"`
	Price     int32              `json:"price" bson:"price"`
	Addons    []Addons           `json:"addons" bson:"addons"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by"`
}

type Addons struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Content   string             `json:"content" bson:"content"`
	DeletedAt bool               `json:"deleted_at" bson:"deleted_at,omitempty"`
	Image     string             `json:"image" bson:"image"`
	Price     int64              `json:"price" bson:"price"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by"`
}
