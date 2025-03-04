package repositories

import (
	"context"
	"server/db"

	dtoGamification "server/dto/gamification"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GamificationRepository encapsulates the logic to access gamification from the data source.
func CreateGamificationRequest(ctx context.Context, gamification dtoGamification.CreateGamificationRequest) (dtoGamification.CreateGamificationRequest, error) {
	collection := db.GetCollection("gamification")
	_, err := collection.InsertOne(ctx, gamification)
	return gamification, err
}

// GetAllGamifications retrieves all gamification data from the data source.
func GetAllgamification(ctx context.Context) ([]dtoGamification.CreateGamificationRequest, error) {
	var gamifications []dtoGamification.CreateGamificationRequest
	collection := db.GetCollection("gamification")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var gamification dtoGamification.CreateGamificationRequest
		if err := cursor.Decode(&gamification); err != nil {
			return nil, err
		}
		gamifications = append(gamifications, gamification)
	}
	return gamifications, cursor.Err()
}

// GetGamificationByID retrieves a gamification data by idUser from the data source.
func GetGamificationByIdUser(ctx context.Context, idUser primitive.ObjectID) (dtoGamification.CreateGamificationRequest, error) {
	var gamification dtoGamification.CreateGamificationRequest
	collection := db.GetCollection("gamification")
	filter := bson.M{"_id_user": idUser}
	err := collection.FindOne(ctx, filter).Decode(&gamification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dtoGamification.CreateGamificationRequest{}, nil
		}
		return dtoGamification.CreateGamificationRequest{}, err
	}
	return gamification, nil
}

func GetGamificationById(ctx context.Context, id primitive.ObjectID) (dtoGamification.CreateGamificationRequest, error) {
	var gamification dtoGamification.CreateGamificationRequest
	collection := db.GetCollection("gamification")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&gamification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dtoGamification.CreateGamificationRequest{}, nil
		}
		return dtoGamification.CreateGamificationRequest{}, err
	}
	return gamification, nil
}

func GetUpdatedGamificationByIdUser(ctx context.Context, idUser primitive.ObjectID) (dtoGamification.UpdatePointGamificationRequest, error) {
	var gamification dtoGamification.UpdatePointGamificationRequest
	collection := db.GetCollection("gamification")
	filter := bson.M{"_id_user": idUser}
	err := collection.FindOne(ctx, filter).Decode(&gamification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dtoGamification.UpdatePointGamificationRequest{}, nil
		}
		return dtoGamification.UpdatePointGamificationRequest{}, err
	}
	return gamification, nil
}

func GetUpdatedGamificationById(ctx context.Context, id primitive.ObjectID) (dtoGamification.UpdatePointGamificationRequest, error) {
	var gamification dtoGamification.UpdatePointGamificationRequest
	collection := db.GetCollection("gamification")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&gamification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dtoGamification.UpdatePointGamificationRequest{}, nil
		}
		return dtoGamification.UpdatePointGamificationRequest{}, err
	}
	return gamification, nil
}

func GetActiveGamificationByIdUser(ctx context.Context, idUser primitive.ObjectID) (dtoGamification.ActiveGamificationRequest, error) {
	var gamification dtoGamification.ActiveGamificationRequest
	collection := db.GetCollection("gamification")
	filter := bson.M{"_id_user": idUser}
	err := collection.FindOne(ctx, filter).Decode(&gamification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dtoGamification.ActiveGamificationRequest{}, nil
		}
		return dtoGamification.ActiveGamificationRequest{}, err
	}
	return gamification, nil
}

func GetActiveGamificationById(ctx context.Context, id primitive.ObjectID) (dtoGamification.ActiveGamificationRequest, error) {
	var gamification dtoGamification.ActiveGamificationRequest
	collection := db.GetCollection("gamification")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&gamification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dtoGamification.ActiveGamificationRequest{}, nil
		}
		return dtoGamification.ActiveGamificationRequest{}, err
	}
	return gamification, nil
}

// UpdateGamification updates the gamification data by idUser.
func UpdateGamification(ctx context.Context, id primitive.ObjectID, idUser primitive.ObjectID, updated dtoGamification.UpdatePointGamificationRequest) (*dtoGamification.UpdatePointGamificationRequest, error) {
	collection := db.GetCollection("gamification")
	var filter bson.M
	if id.IsZero() {
		filter = bson.M{"_id_user": idUser}
	} else {
		filter = bson.M{"_id": id}
	}

	update := bson.M{}
	if updated.Point != 0 {
		update["point"] = updated.Point
	}
	update["updated_at"] = updated.UpdatedAt
	update["updated_by"] = updated.UpdatedBy
	updateResult, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	if updateResult.MatchedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}

	if !id.IsZero() {
		result, err := GetUpdatedGamificationById(ctx, id)
		if err != nil {
			return nil, err
		}
		return &result, nil
	} else {
		result, err := GetUpdatedGamificationByIdUser(ctx, idUser)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
}

// ActiveGamification deletes the gamification data by idUser.
func ActiveGamification(ctx context.Context, id primitive.ObjectID, idUser primitive.ObjectID, updated dtoGamification.ActiveGamificationRequest) (*dtoGamification.ActiveGamificationRequest, error) {
	collection := db.GetCollection("gamification")
	var filter bson.M
	if id.IsZero() {
		filter = bson.M{"_id_user": idUser}
	} else {
		filter = bson.M{"_id": id}
	}

	update := bson.M{}
	update["deleted_at"] = updated.DeletedAt
	update["updated_at"] = updated.UpdatedAt
	update["updated_by"] = updated.UpdatedBy
	updateResult, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	if updateResult.MatchedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}

	if !id.IsZero() {
		result, err := GetActiveGamificationById(ctx, id)
		if err != nil {
			return nil, err
		}
		return &result, nil
	} else {
		result, err := GetActiveGamificationByIdUser(ctx, idUser)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
}

func AddChallenge(ctx context.Context, id primitive.ObjectID, idUser primitive.ObjectID, challenge dtoGamification.Challenges) (dtoGamification.Challenges, error) {
	collection := db.GetCollection("gamification")
	var filter = bson.M{}

	if id.IsZero() {
		filter = bson.M{"_id_user": idUser}
	} else {
		filter = bson.M{"_id": id}
	}
	update := bson.M{
		"$push": bson.M{"challenges": challenge},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return challenge, err
}

func GetChallenges(ctx context.Context, id primitive.ObjectID, idUser primitive.ObjectID) ([]dtoGamification.Challenges, error) {
	var gamification struct {
		Challenges []dtoGamification.Challenges `bson:"challenges"`
	}
	collection := db.GetCollection("gamification")
	var filter = bson.M{}
	if id.IsZero() {
		filter = bson.M{"_id_user": idUser}
	} else {
		filter = bson.M{"_id": id}
	}
	err := collection.FindOne(ctx, filter).Decode(&gamification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return gamification.Challenges, nil
}

func GetChallengesById(ctx context.Context, id primitive.ObjectID) (dtoGamification.Challenges, error) {
	var gamification struct {
		Challenges []dtoGamification.Challenges `bson:"challenges"`
	}
	collection := db.GetCollection("gamification")
	filter := bson.M{"challenges._id": id}

	err := collection.FindOne(ctx, filter).Decode(&gamification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dtoGamification.Challenges{}, nil
		}
		return dtoGamification.Challenges{}, err
	}

	for _, challenge := range gamification.Challenges {
		if challenge.Id == id {
			return challenge, nil
		}
	}
	return dtoGamification.Challenges{}, nil
}

func UpdateChallengesById(ctx context.Context, id primitive.ObjectID, updated dtoGamification.Challenges) (dtoGamification.Challenges, error) {
	collection := db.GetCollection("gamification")
	filter := bson.M{"challenges._id": id}

	update := bson.M{}
	if updated.Name != "" {
		update["challenges.$.name"] = updated.Name
	}
	if updated.Description != "" {
		update["challenges.$.description"] = updated.Description
	}
	if updated.Point != 0 {
		update["challenges.$.point"] = updated.Point
	}
	if updated.Progress != 0 {
		update["challenges.$.progress"] = updated.Progress
	}
	if updated.OnProgress != 0 {
		update["challenges.$.on_progress"] = updated.OnProgress
	}
	if updated.Claim {
		update["challenges.$.claim"] = true
	}
	if updated.Sponsor != "" {
		update["challenges.$.sponsor"] = updated.Sponsor
	}
	if !updated.StartDate.IsZero() {
		update["challenges.$.start_date"] = updated.StartDate
	}
	if !updated.EndDate.IsZero() {
		update["challenges.$.end_date"] = updated.EndDate
	}
	update["challenges.$.updated_at"] = updated.UpdatedAt
	update["challenges.$.updated_by"] = updated.UpdatedBy

	updateResult, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return dtoGamification.Challenges{}, err
	}
	if updateResult.MatchedCount == 0 {
		return dtoGamification.Challenges{}, mongo.ErrNoDocuments
	}

	return GetChallengesById(ctx, id)
}
