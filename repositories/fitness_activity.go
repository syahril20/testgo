package repositories

import (
	"context"
	"errors"
	"server/db"
	"server/dto"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFitnessActivity(ctx context.Context, activityDTO dto.FitnessActivityDTO) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collection := db.GetCollection("fitness_activities")

	idUser, err := convertHexToObjectID(activityDTO.IDUser)
	if err != nil {
		return err
	}

	idFitness, err := convertHexToObjectID(activityDTO.IDFitness)
	if err != nil {
		return err
	}

	activity := models.FitnessActivity{
		ID:        primitive.NewObjectID(),
		IDuser:    idUser,
		IDfitness: idFitness,
		Finished:  activityDTO.Finished,
	}

	_, err = collection.InsertOne(ctx, activity)
	return err
}

func GetFitnessActivityByID(ctx context.Context, id string) (dto.FitnessActivityDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collection := db.GetCollection("fitness_activities")

	objID, err := convertHexToObjectID(id)
	if err != nil {
		return dto.FitnessActivityDTO{}, err
	}

	var activity models.FitnessActivity
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&activity)
	if err != nil {
		return dto.FitnessActivityDTO{}, err
	}

	return mapFitnessActivityToDTO(activity), nil
}

func GetAllFitnessActivities(ctx context.Context) ([]dto.FitnessActivityDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collection := db.GetCollection("fitness_activities")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var activities []dto.FitnessActivityDTO
	for cursor.Next(ctx) {
		var activity models.FitnessActivity
		if err := cursor.Decode(&activity); err != nil {
			return nil, err
		}
		activities = append(activities, mapFitnessActivityToDTO(activity))
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return activities, nil
}

func UpdateFitnessActivity(ctx context.Context, id string, activityDTO dto.FitnessActivityDTO) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collection := db.GetCollection("fitness_activities")

	objID, err := convertHexToObjectID(id)
	if err != nil {
		return err
	}

	idUser, err := convertHexToObjectID(activityDTO.IDUser)
	if err != nil {
		return err
	}

	idFitness, err := convertHexToObjectID(activityDTO.IDFitness)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"_id_user":    idUser,
			"_id_fitness": idFitness,
			"finished":    activityDTO.Finished,
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("no document found to update")
	}

	return nil
}

func DeleteFitnessActivity(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collection := db.GetCollection("fitness_activities")

	objID, err := convertHexToObjectID(id)
	if err != nil {
		return err
	}

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no document found to delete")
	}

	return nil
}

// Helper functions
func convertHexToObjectID(id string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, errors.New("invalid object ID format")
	}
	return objID, nil
}

func mapFitnessActivityToDTO(activity models.FitnessActivity) dto.FitnessActivityDTO {
	return dto.FitnessActivityDTO{
		ID:        activity.ID.Hex(),
		IDUser:    activity.IDuser.Hex(),
		IDFitness: activity.IDfitness.Hex(),
		Finished:  activity.Finished,
	}
}
