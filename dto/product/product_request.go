package productDto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateProductRequest struct {
	Id         primitive.ObjectID        `json:"_id" bson:"_id,omitempty"`
	Name       string                    `json:"name" bson:"name"`
	Content    string                    `json:"content" bson:"content"`
	DeletedAt  *time.Time                `json:"deleted_at" bson:"deleted_at"`
	Image      string                    `json:"image" bson:"image"`
	SubProduct []CreateSubProductRequest `json:"sub_product" bson:"sub_product"`
	CreatedAt  time.Time                 `json:"created_at" bson:"created_at"`
	CreatedBy  string                    `json:"created_by" bson:"created_by"`
	UpdatedAt  time.Time                 `json:"updated_at" bson:"updated_at"`
	UpdatedBy  string                    `json:"updated_by" bson:"updated_by"`
}

type CreateSubProductRequest struct {
	Id        primitive.ObjectID    `json:"_id" bson:"_id,omitempty"`
	IdProduct primitive.ObjectID    `json:"_id_product" bson:"_id_product"`
	Name      string                `json:"name" bson:"name"`
	Content   string                `json:"content" bson:"content"`
	DeletedAt *time.Time            `json:"deleted_at" bson:"deleted_at"`
	Image     string                `json:"image" bson:"image"`
	Price     int32                 `json:"price" bson:"price"`
	Addons    []CreateAddonsRequest `json:"addons" bson:"addons"`
	CreatedAt time.Time             `json:"created_at" bson:"created_at"`
	CreatedBy string                `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time             `json:"updated_at" bson:"updated_at"`
	UpdatedBy string                `json:"updated_by" bson:"updated_by"`
}

type CreateAddonsRequest struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	IdSubProduct primitive.ObjectID `json:"_id_sub_product" bson:"_id_sub_product"`
	Name         string             `json:"name" bson:"name"`
	Content      string             `json:"content" bson:"content"`
	DeletedAt    *time.Time         `json:"deleted_at" bson:"deleted_at"`
	Image        string             `json:"image" bson:"image"`
	Price        int32              `json:"price" bson:"price"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy    string             `json:"created_by" bson:"created_by"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy    string             `json:"updated_by" bson:"updated_by"`
}
