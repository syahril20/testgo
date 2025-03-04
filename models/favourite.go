package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Favourite struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	ArticleID primitive.ObjectID `bson:"article_id" json:"article_id"`
	Count     int                `bson:"count" json:"count"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	CreatedBy string             `bson:"createdBy" json:"createdBy"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
	UpdatedBy string             `bson:"updatedBy" json:"updatedBy"`
	DeletedAt time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"` // Soft delete field
}
