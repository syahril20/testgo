package repositories

import (
	"context"
	"errors"
	"server/db"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateArea(ctx context.Context, area models.Area) error {
	collection := db.GetCollection("areas")
	_, err := collection.InsertOne(ctx, area)
	return err
}

func GetAllAreas(ctx context.Context) ([]models.Area, error) {
	collection := db.GetCollection("areas")
	var areas []models.Area

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var area models.Area
		if err := cursor.Decode(&area); err != nil {
			return nil, err
		}
		areas = append(areas, area)
	}

	return areas, nil
}

func GetAreaByID(ctx context.Context, id primitive.ObjectID) (models.Area, error) {
	collection := db.GetCollection("areas")
	var area models.Area

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&area)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Area{}, errors.New("area not found")
		}
		return models.Area{}, err
	}

	return area, nil
}

func UpdateArea(ctx context.Context, id primitive.ObjectID, updatedArea models.Area) error {
	collection := db.GetCollection("areas")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedArea}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func SoftDeleteArea(ctx context.Context, id primitive.ObjectID, deletedAt primitive.DateTime) error {
	collection := db.GetCollection("areas")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"deleted_at": deletedAt}}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
