package repositories

import (
	"context"
	"server/db"

	dtoMembership "server/dto/membership"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMembership(ctx context.Context, membership dtoMembership.CreateMembershipRequest) (dtoMembership.CreateMembershipRequest, error) {
	collection := db.GetCollection("membership")
	_, err := collection.InsertOne(ctx, membership)
	return membership, err
}

func GetMembershipByID(ctx context.Context, id primitive.ObjectID) (dtoMembership.CreateMembershipRequest, error) {
	var membership dtoMembership.CreateMembershipRequest
	collection := db.GetCollection("membership")
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&membership)
	if err == mongo.ErrNoDocuments {
		return membership, nil
	}
	return membership, err
}

func GetMembershipByIdMember(ctx context.Context, idMember primitive.ObjectID) (*dtoMembership.CreateMembershipRequest, error) {
	var membership dtoMembership.CreateMembershipRequest
	collection := db.GetCollection("membership")
	err := collection.FindOne(ctx, bson.M{"_id_member": idMember}).Decode(&membership)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &membership, err
}

func GetMembershipByIdUser(ctx context.Context, idUser primitive.ObjectID) (*dtoMembership.CreateMembershipRequest, error) {
	var membership dtoMembership.CreateMembershipRequest
	collection := db.GetCollection("membership")
	err := collection.FindOne(ctx, bson.M{"_id_user": idUser}).Decode(&membership)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &membership, err
}

func GetAllMembership(ctx context.Context) ([]dtoMembership.CreateMembershipRequest, error) {
	var memberships []dtoMembership.CreateMembershipRequest
	collection := db.GetCollection("membership")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var membership dtoMembership.CreateMembershipRequest
		if err := cursor.Decode(&membership); err != nil {
			return nil, err
		}
		memberships = append(memberships, membership)
	}
	return memberships, cursor.Err()
}

func GetUpdateMembershipByID(ctx context.Context, id primitive.ObjectID) (*dtoMembership.UpdateMembershipRequest, error) {
	var membership dtoMembership.UpdateMembershipRequest
	collection := db.GetCollection("membership")
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&membership)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &membership, err
}

func UpdateMembershipByID(ctx context.Context, id primitive.ObjectID, updateData dtoMembership.UpdateMembershipRequest) (*dtoMembership.UpdateMembershipRequest, error) {
	collection := db.GetCollection("membership")
	filter := bson.M{"_id": id}

	update := bson.M{}
	if updateData.IdMember != "" {
		update["_id_member"] = updateData.IdMember
	}
	if updateData.IdUser != "" {
		update["_id_user"] = updateData.IdUser
	}
	update["updated_at"] = updateData.UpdatedAt
	update["updated_by"] = updateData.UpdatedBy
	// Add other fields to update as needed

	updateResult, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	if updateResult.MatchedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return GetUpdateMembershipByID(ctx, id)
}

func ActiveMembershipByID(ctx context.Context, id primitive.ObjectID, updateData dtoMembership.ActiveMembershipRequest) (*dtoMembership.ActiveMembershipRequest, error) {
	collection := db.GetCollection("membership")
	filter := bson.M{"_id": id}

	update := bson.M{}
	update["deleted_at"] = updateData.DeletedAt
	update["updated_at"] = updateData.UpdatedAt
	update["updated_by"] = updateData.UpdatedBy

	updateResult, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	if updateResult.MatchedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return GetActiveMembershipByID(ctx, id)
}

func GetActiveMembershipByID(ctx context.Context, id primitive.ObjectID) (*dtoMembership.ActiveMembershipRequest, error) {
	var membership dtoMembership.ActiveMembershipRequest
	collection := db.GetCollection("membership")
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&membership)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &membership, err
}
