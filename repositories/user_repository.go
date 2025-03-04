package repositories

import (
	"context"
	"errors"
	"server/db"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrUserNotFound = errors.New("user not found")
var ErrEmailExists = errors.New("email already exists")

func CreateUser(ctx context.Context, user models.User) error {
	collection := db.GetCollection("users")

	// Check if email already exists
	existingUser := collection.FindOne(ctx, bson.M{"email.value": user.Email.Value})
	if existingUser.Err() == nil {
		return ErrEmailExists
	}

	_, err := collection.InsertOne(ctx, user)
	return err
}

func GetUserByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	collection := db.GetCollection("users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, errors.New("invalid ID format")
	}

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	return user, err
}

func GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	collection := db.GetCollection("users")

	err := collection.FindOne(ctx, bson.M{"email.value": email}).Decode(&user)
	return user, err
}

func GetUserByEmailV2(ctx context.Context, email string) (models.UserResponse, error) {
	var user models.UserResponse
	collection := db.GetCollection("users")

	err := collection.FindOne(ctx, bson.M{"email.value": email}).Decode(&user)
	return user, err
}

func UpdateUser(ctx context.Context, email string, user models.User) error {
	collection := db.GetCollection("users")
	update := bson.M{
		"$set": bson.M{
			"data.name":      user.Data.Name,
			"data.birthDate": user.Data.BirthDate,
			"data.gender":    user.Data.Gender,
			"level":          user.Level,
			"updated_at":     user.UpdatedAt,
			"updated_by":     user.UpdatedBy,
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"email.value": email}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return ErrUserNotFound
	}

	return nil
}
