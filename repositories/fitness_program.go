package repositories

import (
	"context"
	"server/db"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFitnessProgram(ctx context.Context, program models.FitnessProgram) error {
	collection := db.GetCollection("fitness_programs")
	_, err := collection.InsertOne(ctx, program)
	return err
}

func GetAllFitnessPrograms(ctx context.Context) ([]models.FitnessProgram, error) {
	collection := db.GetCollection("fitness_programs")
	cursor, err := collection.Find(ctx, bson.M{"deleted_at": nil}) // Filter: hanya data aktif
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var programs []models.FitnessProgram
	if err := cursor.All(ctx, &programs); err != nil {
		return nil, err
	}

	return programs, nil
}

func GetFitnessProgramByID(ctx context.Context, id string) (models.FitnessProgram, error) {
	collection := db.GetCollection("fitness_programs")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.FitnessProgram{}, err
	}

	var program models.FitnessProgram
	err = collection.FindOne(ctx, bson.M{"_id": objectID, "deleted_at": nil}).Decode(&program) // Data belum dihapus
	return program, err
}

func UpdateFitnessProgram(ctx context.Context, id string, updates bson.M) error {
	collection := db.GetCollection("fitness_programs")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	updates["updated_at"] = time.Now()
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID, "deleted_at": nil}, bson.M{"$set": updates})
	return err
}

func DeleteFitnessProgram(ctx context.Context, id string) error {
	collection := db.GetCollection("fitness_programs")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	now := time.Now()
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": bson.M{"deleted_at": now}})
	return err
}
