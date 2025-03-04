package repositories

import (
	"context"
	"server/db"

	dtoMember "server/dto/member"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMember(ctx context.Context, member dtoMember.CreateMemberRequest) (dtoMember.CreateMemberRequest, error) {
	collection := db.GetCollection("member")
	_, err := collection.InsertOne(ctx, member)
	return member, err
}

func GetMemberByID(ctx context.Context, id primitive.ObjectID) (dtoMember.CreateMemberRequest, error) {
	var member dtoMember.CreateMemberRequest
	collection := db.GetCollection("member")
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&member)
	if err == mongo.ErrNoDocuments {
		return member, nil
	}
	return member, err
}

func GetMemberByName(ctx context.Context, name string) ([]dtoMember.CreateMemberRequest, error) {
	var members []dtoMember.CreateMemberRequest
	collection := db.GetCollection("member")
	cursor, err := collection.Find(ctx, bson.M{"name": name})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var member dtoMember.CreateMemberRequest
		if err := cursor.Decode(&member); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, cursor.Err()
}

func GetAllMembers(ctx context.Context) ([]dtoMember.CreateMemberRequest, error) {
	var members []dtoMember.CreateMemberRequest
	collection := db.GetCollection("member")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var member dtoMember.CreateMemberRequest
		if err := cursor.Decode(&member); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, cursor.Err()
}

func GetActiveMembers(ctx context.Context) ([]dtoMember.CreateMemberRequest, error) {
	members, err := GetAllMembers(ctx)
	if err != nil {
		return nil, err
	}

	var activeMembers []dtoMember.CreateMemberRequest
	for _, member := range members {
		if member.DeletedAt == nil {
			activeMembers = append(activeMembers, member)
		}
	}

	if len(activeMembers) == 0 {
		return nil, nil
	}
	return activeMembers, nil
}

func GetNonActiveMembers(ctx context.Context) ([]dtoMember.CreateMemberRequest, error) {
	members, err := GetAllMembers(ctx)
	if err != nil {
		return nil, err
	}

	var activeMembers []dtoMember.CreateMemberRequest
	for _, member := range members {
		if member.DeletedAt != nil {
			activeMembers = append(activeMembers, member)
		}
	}

	if len(activeMembers) == 0 {
		return nil, nil
	}
	return activeMembers, nil
}

func UpdateMemberByID(ctx context.Context, id primitive.ObjectID, update dtoMember.UpdateMemberRequest) (dtoMember.UpdateMemberRequest, error) {
	// var updatedMember dtoMember.UpdateMemberRequest
	collection := db.GetCollection("member")
	filter := bson.M{"_id": id}
	updateData := bson.M{}
	if update.Name != "" {
		updateData["name"] = update.Name
	}
	if update.Price != 0 {
		updateData["price"] = update.Price
	}
	if update.Benefit != "" {
		updateData["benefit"] = update.Benefit
	}
	updateData = bson.M{"$set": updateData}
	_, err := collection.UpdateOne(ctx, filter, updateData)
	if err == mongo.ErrNoDocuments {
		return update, nil
	}
	return update, err
}

func ActiveMemberByID(ctx context.Context, id primitive.ObjectID, update dtoMember.ActiveMemberRequest) (dtoMember.ActiveMemberRequest, error) {
	// var updatedMember dtoMember.UpdateMemberRequest
	collection := db.GetCollection("member")
	filter := bson.M{"_id": id}
	updateData := bson.M{"$set": update}
	_, err := collection.UpdateOne(ctx, filter, updateData)
	if err == mongo.ErrNoDocuments {
		return update, nil
	}
	return update, err
}
