package repositories

import (
	"context"
	"server/db"
	dto "server/dto/article/article_nutrition"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateArticleNutrition(ctx context.Context, article models.NutritionArticle) error {
	collection := db.GetCollection("article_nutritions")
	_, err := collection.InsertOne(ctx, article)
	return err
}

func UpdateArticleNutrition(ctx context.Context, id primitive.ObjectID, updatedArticle models.NutritionArticle) error {
	collection := db.GetCollection("article_nutritions")
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updatedArticle},
	)
	return err
}

// GetArticleNutritionByID: Mengambil artikel nutrisi berdasarkan ID
func GetArticleNutritionByID(ctx context.Context, id primitive.ObjectID) (models.NutritionArticle, error) {
	collection := db.GetCollection("article_nutritions")
	var article models.NutritionArticle
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&article)
	return article, err
}

func DeleteArticleNutrition(ctx context.Context, id primitive.ObjectID) error {
	collection := db.GetCollection("article_nutrition")

	// Perbarui dokumen dengan menambahkan field `deleted_at`
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"deleted_at": time.Now()}},
	)
	return err
}

func GetAllArticleNutritions(ctx context.Context) ([]dto.NutritionArticleRequest, error) {
	collection := db.GetCollection("article_nutritions")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var articleNutritions []dto.NutritionArticleRequest
	for cursor.Next(ctx) {
		var article models.NutritionArticle
		if err := cursor.Decode(&article); err != nil {
			return nil, err
		}
		articleNutritions = append(articleNutritions, convertArticleNutritionToDTO(article))
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return articleNutritions, nil
}

func convertArticleNutritionToDTO(article models.NutritionArticle) dto.NutritionArticleRequest {
	return dto.NutritionArticleRequest{
		IDCategory:   article.CategoryID,
		Title:        article.Title,
		Image:        article.Image,
		Content:      article.Content,
		TimeToCook:   article.TimeToCook,
		ServingSize:  article.ServingSize,
		Nutritions:   convertNutritionModelToDTO(article.Nutritions),
		Ingredients:  article.Ingredients,
		Instructions: article.Instructions,
		CreatedBy:    article.CreatedBy,
		CreatedAt:    time.Time{}, // Memastikan format sudah benar
		UpdatedBy:    article.UpdatedBy,
		UpdatedAt:    time.Time{}, // Memastikan format sudah benar
	}
}

func convertNutritionModelToDTO(nutritions []models.Nutrition) []dto.Nutrition {
	var nutritionDTOs []dto.Nutrition
	for _, n := range nutritions {
		nutritionDTOs = append(nutritionDTOs, dto.Nutrition{
			Title: n.Title,
			Value: n.Value,
			Unit:  n.Unit,
		})
	}
	return nutritionDTOs
}
