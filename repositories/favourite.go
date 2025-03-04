package repositories

import (
	"context"
	"server/db"
	"server/dto/favourite"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFavorite(ctx context.Context, favorite models.Favourite) error {
	collection := db.GetCollection("favorites")
	_, err := collection.InsertOne(ctx, favorite)
	return err
}

func GetFavoriteByID(ctx context.Context, id primitive.ObjectID) (models.Favourite, error) {
	collection := db.GetCollection("favorites")
	var favourite models.Favourite

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&favourite)
	return favourite, err
}

func GetAllFavorites(ctx context.Context) ([]models.Favourite, error) {
	collection := db.GetCollection("favorites")

	// Menambahkan filter untuk menghindari data yang dihapus
	cursor, err := collection.Find(ctx, bson.M{"deletedAt": nil})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var favorites []models.Favourite
	for cursor.Next(ctx) {
		var favorite models.Favourite
		if err := cursor.Decode(&favorite); err != nil {
			return nil, err
		}
		favorites = append(favorites, favorite)
	}

	return favorites, nil
}

func UpdateFavorite(ctx context.Context, id primitive.ObjectID, favoriteDTO dto.FavouriteUpdateDTO) error {
	collection := db.GetCollection("favorites")

	// Update data yang ada, pastikan tidak ada perubahan jika sudah dihapus
	update := bson.M{
		"$set": bson.M{
			"userID":    favoriteDTO.UserID,
			"articleID": favoriteDTO.ArticleID,
			"count":     favoriteDTO.Count,
			"updatedBy": favoriteDTO.UpdatedBy,
			"updatedAt": time.Now(),
		},
	}

	// Memastikan data belum dihapus dengan filter "deletedAt": nil
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id, "deletedAt": nil}, update)
	return err
}

func DeleteFavorite(ctx context.Context, id primitive.ObjectID, deletedAt time.Time) error {
	collection := db.GetCollection("favorites")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"deletedAt": deletedAt}}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
