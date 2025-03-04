package repositories

import (
	"context"
	"errors"
	"server/db"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateSuburb(ctx context.Context, suburb models.Suburb) error {
	collection := db.GetCollection("suburbs")
	_, err := collection.InsertOne(ctx, suburb)
	return err
}

func GetSuburbByID(ctx context.Context, id string) (models.Suburb, error) {
	var suburb models.Suburb
	collection := db.GetCollection("suburbs")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return suburb, errors.New("invalid ID format")
	}

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&suburb)
	return suburb, err
}

func UpdateSuburb(ctx context.Context, id primitive.ObjectID, suburb models.Suburb) error {
	collection := db.GetCollection("suburbs")

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": suburb},
	)
	return err
}

func GetAllSuburbs(ctx context.Context) ([]models.Suburb, error) {
	collection := db.GetCollection("suburbs")

	var suburbs []models.Suburb
	cursor, err := collection.Find(ctx, bson.M{"deleted_at": nil})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var suburb models.Suburb
		if err := cursor.Decode(&suburb); err != nil {
			return nil, err
		}
		suburbs = append(suburbs, suburb)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return suburbs, nil
}

func DeleteSuburb(ctx context.Context, id primitive.ObjectID, deletedAt time.Time) error {
	collection := db.GetCollection("suburbs")

	// Filter untuk mencari suburb berdasarkan ID
	filter := bson.M{"_id": id}

	// Mengupdate field deleted_at untuk melakukan soft delete
	update := bson.M{"$set": bson.M{"deleted_at": deletedAt}}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
