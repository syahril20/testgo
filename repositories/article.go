package repositories

import (
	"context"
	"server/db"
	"server/dto/article"
	"server/models"
	"server/pkg" // Import pkg untuk ParseObjectID

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateArticle: Menambahkan artikel baru
func CreateArticle(ctx context.Context, articleDTO dto.Article) error {
	collection := db.GetCollection("articles")

	// Konversi DTO ke Model
	idCategory, err := pkg.ParseObjectID(articleDTO.IDCategory)
	if err != nil {
		return err
	}

	article := models.Article{
		ID:          primitive.NewObjectID(),
		IDCategory:  idCategory,
		Title:       articleDTO.Title,
		Premium:     articleDTO.Premium,
		ReadingTime: articleDTO.ReadingTime,
		Image:       articleDTO.Image,
		Detail: models.DetailArticle{
			Description:   articleDTO.Detail.Description,
			Ingredients:   articleDTO.Detail.Ingredients,
			Instructions:  articleDTO.Detail.Instructions,
			CountTime:     articleDTO.Detail.CountTime,
			CountCalories: articleDTO.Detail.CountCalories,
		},
		Nutrition: parseNutrition(articleDTO.Nutrition),
		Index:     articleDTO.Index,
		Deleted:   articleDTO.Deleted,
	}

	_, err = collection.InsertOne(ctx, article)
	return err
}

// Helper function: Convert NutritionItems
func parseNutrition(nutritionDTO []dto.NutritionItem) []models.NutritionItem {
	var nutrition []models.NutritionItem
	for _, item := range nutritionDTO {
		nutrition = append(nutrition, models.NutritionItem{
			Type:  item.Type,
			Icon:  item.Icon,
			Value: item.Value,
		})
	}
	return nutrition
}

// GetArticleByID: Mengambil artikel berdasarkan ID
func GetArticleByID(ctx context.Context, id string) (dto.Article, error) {
	collection := db.GetCollection("articles")

	objID, err := pkg.ParseObjectID(id)
	if err != nil {
		return dto.Article{}, err
	}

	var article models.Article
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&article)
	if err != nil {
		return dto.Article{}, err
	}

	// Konversi Model ke DTO
	return convertArticleToDTO(article), nil
}

// Helper function: Konversi Model ke DTO
func convertArticleToDTO(article models.Article) dto.Article {
	nutritionDTO := []dto.NutritionItem{}
	for _, item := range article.Nutrition {
		nutritionDTO = append(nutritionDTO, dto.NutritionItem{
			Type:  item.Type,
			Icon:  item.Icon,
			Value: item.Value,
		})
	}

	return dto.Article{
		ID:          article.ID.Hex(),
		IDCategory:  article.IDCategory.Hex(),
		Title:       article.Title,
		Premium:     article.Premium,
		ReadingTime: article.ReadingTime,
		Image:       article.Image,
		Detail: dto.DetailArticle{
			Description:   article.Detail.Description,
			Ingredients:   article.Detail.Ingredients,
			Instructions:  article.Detail.Instructions,
			CountTime:     article.Detail.CountTime,
			CountCalories: article.Detail.CountCalories,
		},
		Nutrition: nutritionDTO,
		Index:     article.Index,
		Deleted:   article.Deleted,
	}
}

// GetAllArticles: Mengambil semua artikel
func GetAllArticles(ctx context.Context) ([]dto.Article, error) {
	collection := db.GetCollection("articles")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var articles []dto.Article
	for cursor.Next(ctx) {
		var article models.Article
		if err := cursor.Decode(&article); err != nil {
			return nil, err
		}
		articles = append(articles, convertArticleToDTO(article))
	}

	return articles, nil
}

// DeleteArticle: Menghapus artikel berdasarkan ID
func DeleteArticle(ctx context.Context, id string) error {
	collection := db.GetCollection("articles")

	objID, err := pkg.ParseObjectID(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

// UpdateArticle: Memperbarui artikel berdasarkan ID
func UpdateArticle(ctx context.Context, id string, articleDTO dto.Article) error {
	collection := db.GetCollection("articles")

	objID, err := pkg.ParseObjectID(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":        articleDTO.Title,
			"premium":      articleDTO.Premium,
			"reading_time": articleDTO.ReadingTime,
			"image":        articleDTO.Image,
			"detail": bson.M{
				"description":    articleDTO.Detail.Description,
				"ingredients":    articleDTO.Detail.Ingredients,
				"instructions":   articleDTO.Detail.Instructions,
				"count_time":     articleDTO.Detail.CountTime,
				"count_calories": articleDTO.Detail.CountCalories,
			},
			"nutrition": articleDTO.Nutrition,
			"index":     articleDTO.Index,
			"deleted":   articleDTO.Deleted,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}
