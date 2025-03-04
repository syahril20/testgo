package dto

type Payment struct {
	UserID    string `json:"_id_user" bson:"_id_user"`
	ProductID string `json:"_id_product" bson:"_id_product"`
}

type PaymentRespon struct {
	ID        string `json:"_id" bson:"_id"`
	UserID    string `json:"_id_user" bson:"_id_user"`
	ProductID string `json:"_id_product" bson:"_id_product"`
}
