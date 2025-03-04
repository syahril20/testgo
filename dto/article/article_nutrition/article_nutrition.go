package dto

import "time"

type Nutrition struct {
	Title string `json:"title"`
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type NutritionArticleRequest struct {
	ID           string      `json:"_id"`
	IDCategory   string      `json:"_id_category"`
	Title        string      `json:"title"`
	Image        string      `json:"image"`
	Content      string      `json:"content"`
	TimeToCook   string      `json:"time_to_cook"`
	ServingSize  string      `json:"serving_size"`
	Nutritions   []Nutrition `json:"nutritions"`
	Ingredients  string      `json:"ingredients"`
	Instructions string      `json:"instructions"`
	CreatedAt    time.Time   `json:"created_at" bson:"created_at"`
	CreatedBy    string      `json:"created_by" bson:"created_by"`
	UpdatedAt    time.Time   `json:"updated_at" bson:"updated_at"`
	UpdatedBy    string      `json:"updated_by" bson:"updated_by"`
	DeletedAt    *time.Time  `bson:"deleted_at,omitempty"`
}

type ArticleNutritionResponse struct {
	ID           string      `json:"_id"`
	IDCategory   string      `json:"_id_category"`
	Title        string      `json:"title"`
	Image        string      `json:"image"`
	Content      string      `json:"content"`
	TimeToCook   string      `json:"time_to_cook"`
	ServingSize  string      `json:"serving_size"`
	Nutritions   []Nutrition `json:"nutritions"`
	Ingredients  string      `json:"ingredients"`
	Instructions string      `json:"instructions"`
}
