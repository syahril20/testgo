package dto

type SuburbRespon struct {
	CityID string `json:"_id_city" bson:"_id_city"`
	Name   string `json:"name" bson:"name"`
}
