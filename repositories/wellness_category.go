package repositories

import (
	"context"
	"server/db"
	dto "server/dto/wellness_category"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateWellnessCategory(ctx context.Context, WellnessCategory dto.CreateWellnessCategoryRequest) (dto.CreateWellnessCategoryRequest, error) {
	collection := db.GetCollection("wellness_category")
	_, err := collection.InsertOne(ctx, WellnessCategory)
	return WellnessCategory, err
}

func GetWellnessCategoryByName(ctx context.Context, title string) (*dto.CreateWellnessCategoryRequest, error) {
	collection := db.GetCollection("wellness_category")
	var WellnessCategory dto.CreateWellnessCategoryRequest
	err := collection.FindOne(ctx, bson.M{"title": title}).Decode(&WellnessCategory)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &WellnessCategory, nil
}
