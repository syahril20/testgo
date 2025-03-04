package membership

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateMembershipRequest struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	IdMember  primitive.ObjectID `json:"_id_member" bson:"_id_member"`
	IdUser    primitive.ObjectID `json:"_id_user" bson:"_id_user"`
	EndDate   time.Time          `json:"end_date" bson:"end_date"`
	DeletedAt *time.Time         `json:"deleted_at" bson:"deleted_at"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by"`
}

type UpdateMembershipRequest struct {
	IdMember  string    `json:"_id_member" bson:"_id_member"`
	IdUser    string    `json:"_id_user" bson:"_id_user"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type ActiveMembershipRequest struct {
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
}
