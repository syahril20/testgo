package repositories

import (
	"context"
	"server/db"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentRepository struct{}

func (r *PaymentRepository) CreatePayment(payment models.Payment) error {
	collection := db.GetCollection("payments")
	_, err := collection.InsertOne(context.Background(), payment)
	return err
}

func (r *PaymentRepository) GetPaymentByID(id string) (models.Payment, error) {
	collection := db.GetCollection("payments")
	var payment models.Payment
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return payment, err
	}
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&payment)
	return payment, err
}

func (r *PaymentRepository) UpdatePayment(id string, payment models.Payment) error {
	collection := db.GetCollection("payments")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": payment})
	return err
}

func (r *PaymentRepository) DeletePayment(id string) error {
	collection := db.GetCollection("payments")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}
