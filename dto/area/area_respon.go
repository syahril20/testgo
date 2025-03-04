package dto

type AreaRespon struct {
	ID       string `json:"_id" bson:"_id"`
	SuburbID string `json:"_id_suburb" bson:"_id_suburb"`
	Name     string `json:"name" bson:"name"`
	Zip      string `json:"zip" bson:"zip"`
}
