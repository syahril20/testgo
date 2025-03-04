package repositories

import (
	"context"
	"server/db"
	"server/dto"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateIMS(ctx context.Context, ims models.IMS) error {
	collection := db.GetCollection("ims")
	_, err := collection.InsertOne(ctx, ims)
	return err
}

func GetAllIMS(ctx context.Context) ([]models.IMS, error) {
	collection := db.GetCollection("ims")
	filter := bson.M{"deleted_at": nil}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var imsList []models.IMS
	if err := cursor.All(ctx, &imsList); err != nil {
		return nil, err
	}
	return imsList, nil
}

func GetIMSByID(ctx context.Context, id primitive.ObjectID) (models.IMS, error) {
	collection := db.GetCollection("ims")
	var ims models.IMS

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&ims)
	return ims, err
}

func UpdateIMS(ctx context.Context, id primitive.ObjectID, imsDTO dto.IMSRequest) error {
	collection := db.GetCollection("ims")
	update := bson.M{
		"$set": bson.M{
			"email":                  imsDTO.Email,
			"name":                   imsDTO.Name,
			"old":                    imsDTO.Old,
			"phone":                  imsDTO.Phone,
			"address":                imsDTO.Address,
			"opti_sample_collection": imsDTO.OptiSampleCollection,
			"updated_at":             time.Now(),
			"updated_by":             imsDTO.UpdatedBy,
		},
	}
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id, "deleted_at": nil}, update)
	return err
}

func DeleteIMS(ctx context.Context, id primitive.ObjectID, deletedAt time.Time) error {
	collection := db.GetCollection("ims")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"deletedAt": deletedAt}}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
