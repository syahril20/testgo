package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CityRespon struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ProvinceID string             `json:"_id_province" bson:"_id_province"`
	Name       string             `json:"name" bson:"name"`
}
