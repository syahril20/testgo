package repositories

import (
	"context"
	"server/db"
	dtoAuth "server/dto/auth"
)

func InsertAuthLog(ctx context.Context, authLog dtoAuth.AuthLog) error {
	collection := db.GetCollection("auth_log")
	_, err := collection.InsertOne(ctx, authLog)
	return err
}
