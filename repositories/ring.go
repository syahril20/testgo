package repositories

import (
	"context"
	"server/db"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateRing(ctx context.Context, ring models.Ring) error {
	collection := db.GetCollection("rings")

	_, err := collection.InsertOne(ctx, ring)
	return err
}

func GetRingByID(ctx context.Context, id string) (models.Ring, error) {
	collection := db.GetCollection("rings")

	var ring models.Ring
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&ring)
	return ring, err
}

func UpdateRing(ctx context.Context, id string, updatedRing models.Ring) error {
	collection := db.GetCollection("rings")

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"size":  updatedRing.Size,
			"color": updatedRing.Color,
		},
	})
	return err
}

func DeleteRing(ctx context.Context, id string) error {
	collection := db.GetCollection("rings")

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
