package repositories

import (
	"context"
	"server/db"
	"server/dto"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFitness(ctx context.Context, fitnessDTO dto.FitnessDTO) (dto.FitnessDTO, error) {
	collection := db.GetCollection("fitness")
	categoryID, err := primitive.ObjectIDFromHex(fitnessDTO.CategoryId)
	if err != nil {
		return dto.FitnessDTO{}, err
	}

	// Menambahkan waktu pembuatan dan siapa yang membuat
	now := time.Now()

	// Membuat fitness menggunakan data dari fitnessDTO
	fitness := bson.M{
		"categoryId":  categoryID,
		"title":       fitnessDTO.Title,
		"image":       fitnessDTO.Image,
		"description": fitnessDTO.Description,
		"video":       fitnessDTO.Video,
		"workout":     fitnessDTO.Workout,
		"created_at":  now,                  // Waktu dibuat
		"created_by":  fitnessDTO.CreatedBy, // Siapa yang membuat
	}

	// Menyimpan fitness ke dalam database
	result, err := collection.InsertOne(ctx, fitness)
	if err != nil {
		return dto.FitnessDTO{}, err
	}

	// Mengambil InsertedID dan mengonversinya menjadi string
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		fitnessDTO.ID = oid.Hex() // Mengonversi ObjectID ke string
	}

	// Mengembalikan fitnessDTO dengan ID yang sudah terisi
	return fitnessDTO, nil
}

func GetFitnessByID(ctx context.Context, id string) (dto.FitnessDTO, error) {
	collection := db.GetCollection("fitness")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.FitnessDTO{}, err
	}

	var fitness dto.FitnessDTO
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&fitness)
	if err != nil {
		return dto.FitnessDTO{}, err
	}

	return fitness, nil
}

func GetAllFitness(ctx context.Context) ([]models.Fitness, error) {
	collection := db.GetCollection("fitness")

	// Melakukan pencarian ke database MongoDB untuk semua dokumen dalam koleksi fitness
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Menyimpan hasil query ke dalam slice
	var fitnessList []models.Fitness
	for cursor.Next(ctx) {
		var fitness models.Fitness
		// Mendekode dokumen MongoDB ke dalam model Fitness
		if err := cursor.Decode(&fitness); err != nil {
			return nil, err
		}

		// Assign ID langsung sebagai string jika menggunakan tipe string
		fitness.ID = fitness.ID

		// Menambahkan fitness ke dalam list
		fitnessList = append(fitnessList, fitness)
	}

	// Mengecek apakah ada error pada cursor
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return fitnessList, nil
}

func UpdateFitness(ctx context.Context, id string, fitnessDTO dto.FitnessDTO) error {
	collection := db.GetCollection("fitness")

	categoryID, err := primitive.ObjectIDFromHex(fitnessDTO.CategoryId)
	if err != nil {
		return err
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	now := time.Now()

	update := bson.M{
		"$set": bson.M{
			"title":       fitnessDTO.Title,
			"image":       fitnessDTO.Image,
			"description": fitnessDTO.Description,
			"categoryId":  categoryID,
			"video":       fitnessDTO.Video,
			"workout":     fitnessDTO.Workout,
			"updated_at":  now, // Waktu dibuat
			"updated_by":  fitnessDTO.CreatedBy,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func DeleteFitness(ctx context.Context, id string) error {
	collection := db.GetCollection("fitness")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
