package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Nutrition struct {
	Title string `json:"title"`
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type NutritionArticle struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CategoryID   string             `json:"_id_category" bson:"category_id"`
	Title        string             `json:"title" bson:"title"`
	Image        string             `json:"image" bson:"image"`
	Content      string             `json:"content" bson:"content"`
	TimeToCook   string             `json:"time_to_cook" bson:"time_to_cook"`
	ServingSize  string             `json:"serving_size" bson:"serving_size"`
	Nutritions   []Nutrition        `json:"nutritions" bson:"nutritions"`
	Ingredients  string             `json:"ingredients" bson:"ingredients"`
	Instructions string             `json:"instructions" bson:"instructions"`
	CreatedBy    string             `json:"created_by" bson:"created_by"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedBy    string             `json:"updated_by" bson:"updated_by"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt    *time.Time         `bson:"deleted_at,omitempty"`
}
