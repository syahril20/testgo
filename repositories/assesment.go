package repositories

import (
	"context"
	"server/db"
	dtoassessment "server/dto/assessment"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// assessment
func Createassessment(ctx context.Context, assessment dtoassessment.CreateassessmentRequest) (dtoassessment.CreateassessmentRequest, error) {
	collection := db.GetCollection("assessment")
	_, err := collection.InsertOne(ctx, assessment)
	return assessment, err
}

func GetassessmentByName(ctx context.Context, name string) (*dtoassessment.CreateassessmentRequest, error) {
	collection := db.GetCollection("assessment")
	var product dtoassessment.CreateassessmentRequest
	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func GetassessmentById(ctx context.Context, id primitive.ObjectID) (*dtoassessment.CreateassessmentRequest, error) {
	collection := db.GetCollection("assessment")
	var assessment dtoassessment.CreateassessmentRequest
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&assessment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &assessment, nil
}

func GetActiveassessmentById(ctx context.Context, id primitive.ObjectID) (*dtoassessment.CreateassessmentRequest, error) {
	collection := db.GetCollection("assessment")
	var assessment dtoassessment.CreateassessmentRequest

	filter := bson.M{"_id": id, "deleted_at": nil}
	err := collection.FindOne(ctx, filter).Decode(&assessment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &assessment, nil
}

func GetAllassessments(ctx context.Context) ([]dtoassessment.CreateassessmentRequest, error) {
	collection := db.GetCollection("assessment")
	var assessments []dtoassessment.CreateassessmentRequest

	filter := bson.M{"deleted_at": bson.M{"$eq": nil}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var assessment dtoassessment.CreateassessmentRequest
		if err := cursor.Decode(&assessment); err != nil {
			return nil, err
		}
		assessments = append(assessments, assessment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return assessments, nil
}

func GetAllNonActiveassessments(ctx context.Context) ([]dtoassessment.CreateassessmentRequest, error) {
	collection := db.GetCollection("assessment")
	var assessments []dtoassessment.CreateassessmentRequest

	filter := bson.M{"deleted_at": bson.M{"$ne": nil}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var assessment dtoassessment.CreateassessmentRequest
		if err := cursor.Decode(&assessment); err != nil {
			return nil, err
		}
		assessments = append(assessments, assessment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return assessments, nil
}

func UpdateassessmentById(ctx context.Context, id primitive.ObjectID, updateData dtoassessment.UpdateassessmentRequest) (*dtoassessment.UpdateassessmentRequest, error) {
	collection := db.GetCollection("assessment")
	var updated dtoassessment.UpdateassessmentRequest

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": updateData,
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, filter).Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func ActiveassessmentById(ctx context.Context, id primitive.ObjectID, updateData dtoassessment.ActiveassessmentRequest) (*dtoassessment.ActiveassessmentRequest, error) {
	collection := db.GetCollection("assessment")
	var updated dtoassessment.ActiveassessmentRequest

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": updateData,
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, filter).Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

// Questionnaire

func CreateQuestionnaire(ctx context.Context, assessmentId primitive.ObjectID, questionnaire dtoassessment.CreateQuestionnaireRequest) (dtoassessment.CreateQuestionnaireRequest, error) {
	collection := db.GetCollection("assessment")
	filter := bson.M{"_id": assessmentId}
	update := bson.M{
		"$push": bson.M{"questionnaire": questionnaire},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return questionnaire, err
}

func GetQuestionnaireByQuestion(ctx context.Context, question string) (*dtoassessment.CreateQuestionnaireRequest, error) {
	collection := db.GetCollection("assessment")
	var questionnaire dtoassessment.CreateQuestionnaireRequest

	filter := bson.M{
		"questionnaire.question": question,
	}

	err := collection.FindOne(ctx, filter).Decode(&questionnaire)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &questionnaire, nil
}

func GetAllDataQuestionnaires(ctx context.Context, assessmentId primitive.ObjectID) ([]dtoassessment.CreateQuestionnaireRequest, error) {
	collection := db.GetCollection("assessment")
	var result []dtoassessment.CreateQuestionnaireRequest

	filter := bson.M{
		"_id":        assessmentId,
		"deleted_at": bson.M{"$eq": nil},
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var assessment struct {
			Questionnaire []dtoassessment.CreateQuestionnaireRequest `bson:"questionnaire"`
		}

		if err := cursor.Decode(&assessment); err != nil {
			return nil, err
		}
		for _, questionnaire := range assessment.Questionnaire {
			if questionnaire.DeletedAt == nil {
				result = append(result, questionnaire)
			}
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetAllNonActiveQuestionnaires(ctx context.Context, assessmentId primitive.ObjectID) ([]dtoassessment.CreateQuestionnaireRequest, error) {
	collection := db.GetCollection("assessment")
	var result []dtoassessment.CreateQuestionnaireRequest

	filter := bson.M{
		"_id":        assessmentId,
		"deleted_at": bson.M{"$eq": nil},
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var assessment struct {
			Questionnaire []dtoassessment.CreateQuestionnaireRequest `bson:"questionnaire"`
		}

		if err := cursor.Decode(&assessment); err != nil {
			return nil, err
		}
		for _, questionnaire := range assessment.Questionnaire {
			if questionnaire.DeletedAt != nil {
				result = append(result, questionnaire)
			}
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetQuestionnaireById(ctx context.Context, id primitive.ObjectID) (*dtoassessment.CreateQuestionnaireRequest, error) {
	collection := db.GetCollection("assessment")
	var questionnaire dtoassessment.CreateassessmentRequest

	filter := bson.M{"questionnaire._id": id}

	err := collection.FindOne(ctx, filter).Decode(&questionnaire)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	if len(questionnaire.Questionnaire) > 0 {
		return &questionnaire.Questionnaire[0], nil
	}
	return nil, nil
}

func UpdateQuestionnaireById(ctx context.Context, id primitive.ObjectID, updateData dtoassessment.UpdateQuestionnaireRequest) (*dtoassessment.UpdateQuestionnaireRequest, error) {
	collection := db.GetCollection("assessment")
	var updated dtoassessment.UpdateQuestionnaireRequest

	filter := bson.M{"questionnaire._id": id}

	update := bson.M{
		"$set": bson.M{
			"questionnaire.$.question":   updateData.Question,
			"questionnaire.$.updated_by": updateData.UpdatedBy,
			"questionnaire.$.updated_at": updateData.UpdatedAt,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, filter).Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func ActiveQuestionnaireById(ctx context.Context, id primitive.ObjectID, updateData dtoassessment.ActiveQuestionnaireRequest) (*dtoassessment.ActiveQuestionnaireRequest, error) {
	collection := db.GetCollection("assessment")
	var updated dtoassessment.ActiveQuestionnaireRequest

	filter := bson.M{"questionnaire._id": id}

	update := bson.M{
		"$set": bson.M{
			"questionnaire.$.updated_by": updateData.UpdatedBy,
			"questionnaire.$.updated_at": updateData.UpdatedAt,
			"questionnaire.$.deleted_at": updateData.DeletedAt,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, filter).Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}
