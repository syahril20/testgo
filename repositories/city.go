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

func CreateCity(ctx context.Context, city models.City) error {
	collection := db.GetCollection("cities")
	_, err := collection.InsertOne(ctx, city)
	return err
}

func GetAllCities(ctx context.Context) ([]models.City, error) {
	collection := db.GetCollection("cities")

	filter := bson.M{"deleted_at": bson.M{"$exists": false}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var cities []models.City
	for cursor.Next(ctx) {
		var city models.City
		if err := cursor.Decode(&city); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	return cities, nil
}

func GetCityByID(ctx context.Context, id string) (models.City, error) {
	var city models.City
	collection := db.GetCollection("cities")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return city, errors.New("invalid ID format")
	}

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&city)
	return city, err
}

func UpdateCity(ctx context.Context, id primitive.ObjectID, city models.City) error {
	collection := db.GetCollection("cities")

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": city},
	)
	return err
}

func SoftDeleteCity(ctx context.Context, id primitive.ObjectID, deletedAt time.Time) error {
	collection := db.GetCollection("cities")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"deleted_at": deletedAt}}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
