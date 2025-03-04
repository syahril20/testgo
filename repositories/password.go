package repositories

import (
	"context"
	"fmt"
	"server/db"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPasswordByID(ctx context.Context, userID string) (models.UserPassword, error) {
	collection := db.GetCollection("users")

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.UserPassword{}, err
	}

	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return models.UserPassword{}, err
	}

	return user.Password, nil
}

func UpdatePassword(ctx context.Context, userID string, newPassword string) error {
	collection := db.GetCollection("users")

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		fmt.Printf("Invalid userID: %s, error: %v\n", userID, err)
		return err
	}

	fmt.Printf("Updating password for userID: %s with new password hash: %s\n", userID, newPassword)

	update := bson.M{
		"$set": bson.M{
			"password.value":          newPassword,
			"password.request_change": false,
			"password.updated_at":     time.Now(),
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		fmt.Printf("Failed to update password for userID %s: %v\n", userID, err)
		return err
	}

	if result.MatchedCount == 0 {
		fmt.Printf("No document found for userID %s\n", userID)
		return fmt.Errorf("user not found")
	}

	fmt.Printf("Password updated successfully for userID: %s\n", userID)
	return nil
}

func RequestPasswordReset(ctx context.Context, userID string, token string, expireTime time.Time) error {
	collection := db.GetCollection("users")

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{
		"$set": bson.M{
			"password.request_forgot": true,
			"password.token":          token,
			"password.request_expire": expireTime,
		},
	})
	return err
}

func ConfirmPasswordReset(ctx context.Context, userID string, newPassword string) error {
	collection := db.GetCollection("users")

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{
		"$set": bson.M{
			"password.value":          newPassword,
			"password.request_forgot": false,
			"password.token":          nil,
			"password.request_expire": nil,
		},
	})
	return err
}
