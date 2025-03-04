package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DetailArticle berisi detail dari artikel
type DetailArticle struct {
	Description   string `bson:"description"`
	Ingredients   string `bson:"ingredients"`
	Instructions  string `bson:"instructions"`
	CountTime     string `bson:"count_time"`
	CountCalories string `bson:"count_calories"`
}

// NutritionItem berisi informasi tentang nutrisi dalam artikel
type NutritionItem struct {
	Type  string `bson:"type"`
	Icon  string `bson:"icon"`
	Value string `bson:"value"`
}

// Article adalah struktur data untuk koleksi artikel
type Article struct {
	ID          primitive.ObjectID `bson:"_id"`
	IDCategory  primitive.ObjectID `bson:"_id_category"`
	Title       string             `bson:"title"`
	Premium     bool               `bson:"premium"`
	ReadingTime int                `bson:"reading_time"`
	Image       string             `bson:"image"`
	Detail      DetailArticle      `bson:"detail"`
	Nutrition   []NutritionItem    `bson:"nutrition"`
	Index       string             `bson:"_index"`
	Deleted     bool               `bson:"deleted"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy   string             `json:"updated_by" bson:"updated_by"`
}
