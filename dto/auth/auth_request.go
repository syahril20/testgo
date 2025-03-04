package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type AuthLog struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	IdUser    primitive.ObjectID `bson:"_id_user" json:"_id_user"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
}
