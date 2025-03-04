package repositories

import (
	"context"
	"errors"
	"server/db"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProvince(ctx context.Context, province models.Province) error {
	collection := db.GetCollection("provinces")
	_, err := collection.InsertOne(ctx, province)
	return err
}

func GetAllProvinces(ctx context.Context) ([]models.Province, error) {
	collection := db.GetCollection("provinces")

	filter := bson.M{"deleted_at": bson.M{"$exists": false}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var provinces []models.Province
	for cursor.Next(ctx) {
		var province models.Province
		if err := cursor.Decode(&province); err != nil {
			return nil, err
		}
		provinces = append(provinces, province)
	}

	return provinces, nil
}

func GetProvinceByID(ctx context.Context, id string) (models.Province, error) {
	var province models.Province
	collection := db.GetCollection("provinces")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return province, errors.New("invalid ID format")
	}

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&province)
	return province, err
}

func UpdateProvince(ctx context.Context, id primitive.ObjectID, province models.Province) error {
	collection := db.GetCollection("provinces")

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": province},
	)
	return err
}

func SoftDeleteProvince(ctx context.Context, id primitive.ObjectID, deletedAt time.Time) error {
	collection := db.GetCollection("provinces")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"deleted_at": deletedAt}}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
