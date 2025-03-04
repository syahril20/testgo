package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/db"
	"server/dto"
	"server/models"
)

func CreateFitnessCategory(ctx context.Context, categoryDTO dto.FitnessCategoryDTO) error {
	collection := db.GetCollection("fitness_categories")

	category := models.FitnessCategory{
		ID:          primitive.NewObjectID(),
		Title:       categoryDTO.Title,
		Description: categoryDTO.Description,
		Image:       categoryDTO.Image,
		Deleted:     categoryDTO.Deleted,
	}

	_, err := collection.InsertOne(ctx, category)
	return err
}

func UpdateFitnessCategory(ctx context.Context, id string, categoryDTO dto.FitnessCategoryDTO) error {
	collection := db.GetCollection("fitness_categories")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":       categoryDTO.Title,
			"description": categoryDTO.Description,
			"image":       categoryDTO.Image,
			"deleted":     categoryDTO.Deleted,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func DeleteFitnessCategory(ctx context.Context, id string) error {
	collection := db.GetCollection("fitness_categories")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func GetFitnessCategoryByID(ctx context.Context, id string) (models.FitnessCategory, error) {
	collection := db.GetCollection("fitness_categories")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.FitnessCategory{}, err
	}

	var category models.FitnessCategory
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&category)
	if err != nil {
		return models.FitnessCategory{}, err
	}

	return category, nil
}

func GetAllFitnessCategories(ctx context.Context) ([]models.FitnessCategory, error) {
	collection := db.GetCollection("fitness_categories")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []models.FitnessCategory
	for cursor.Next(ctx) {
		var category models.FitnessCategory
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
