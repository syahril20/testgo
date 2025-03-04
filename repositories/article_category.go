package repositories

import (
	"context"
	"server/db"
	"server/dto/article/article_category"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory(ctx context.Context, categoryDTO dto.CategoryRequest) error {
	collection := db.GetCollection("categories")

	category := models.Category{
		ID:    primitive.NewObjectID(),
		Title: categoryDTO.Title,
		Image: categoryDTO.Image,
		Icon:  categoryDTO.Icon,
	}

	_, err := collection.InsertOne(ctx, category)
	return err
}

func GetCategoryByID(ctx context.Context, id string) (dto.CategoryRequest, error) {
	collection := db.GetCollection("categories")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.CategoryRequest{}, err
	}

	var category models.Category
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&category)
	if err != nil {
		return dto.CategoryRequest{}, err
	}

	return convertCategoryToDTO(category), nil
}

func UpdateCategory(ctx context.Context, id string, categoryDTO dto.CategoryRequest) error {
	collection := db.GetCollection("categories")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title": categoryDTO.Title,
			"image": categoryDTO.Image,
			"icon":  categoryDTO.Icon,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func DeleteCategory(ctx context.Context, id string) error {
	collection := db.GetCollection("categories")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func GetAllCategories(ctx context.Context) ([]dto.CategoryRequest, error) {
	collection := db.GetCollection("categories")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []dto.CategoryRequest
	for cursor.Next(ctx) {
		var category models.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, convertCategoryToDTO(category))
	}

	return categories, nil
}

func convertCategoryToDTO(category models.Category) dto.CategoryRequest {
	return dto.CategoryRequest{
		ID:    category.ID.Hex(),
		Title: category.Title,
		Image: category.Image,
		Icon:  category.Icon,
	}
}
