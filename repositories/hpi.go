package repositories

import (
	"context"
	"server/db"

	dtoHpi "server/dto/hpi"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// HPI

func CreateHPI(ctx context.Context, hpi dtoHpi.CreateHpiRequest) (dtoHpi.CreateHpiRequest, error) {
	collection := db.GetCollection("hpi")
	_, err := collection.InsertOne(ctx, hpi)
	return hpi, err
}

func GetHPIByName(ctx context.Context, name string) (*dtoHpi.CreateHpiRequest, error) {
	collection := db.GetCollection("hpi")
	var hpi dtoHpi.CreateHpiRequest
	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&hpi)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &hpi, nil
}

func GetHPIById(ctx context.Context, id primitive.ObjectID) (*dtoHpi.CreateHpiRequest, error) {
	collection := db.GetCollection("hpi")
	var hpi dtoHpi.CreateHpiRequest
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&hpi)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &hpi, nil
}

func GetActiveHPIById(ctx context.Context, id primitive.ObjectID) (*dtoHpi.CreateHpiRequest, error) {
	collection := db.GetCollection("hpi")
	var hpi dtoHpi.CreateHpiRequest

	filter := bson.M{"_id": id, "deleted_at": nil}
	err := collection.FindOne(ctx, filter).Decode(&hpi)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &hpi, nil
}

func GetAllHPIs(ctx context.Context) ([]dtoHpi.CreateHpiRequest, error) {
	collection := db.GetCollection("hpi")
	var hpis []dtoHpi.CreateHpiRequest

	filter := bson.M{"deleted_at": bson.M{"$eq": nil}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var hpi dtoHpi.CreateHpiRequest
		if err := cursor.Decode(&hpi); err != nil {
			return nil, err
		}
		hpis = append(hpis, hpi)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return hpis, nil
}

func GetAllNonActiveHPIs(ctx context.Context) ([]dtoHpi.CreateHpiRequest, error) {
	collection := db.GetCollection("hpi")
	var hpis []dtoHpi.CreateHpiRequest

	filter := bson.M{"deleted_at": bson.M{"$ne": nil}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var hpi dtoHpi.CreateHpiRequest
		if err := cursor.Decode(&hpi); err != nil {
			return nil, err
		}
		hpis = append(hpis, hpi)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return hpis, nil
}

func UpdateHPIById(ctx context.Context, id primitive.ObjectID, updateData dtoHpi.UpdateHpiRequest) (*dtoHpi.UpdateHpiRequest, error) {
	collection := db.GetCollection("hpi")
	var updated dtoHpi.UpdateHpiRequest

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

func ActiveHPIById(ctx context.Context, id primitive.ObjectID, updateData dtoHpi.ActiveHpiRequest) (*dtoHpi.ActiveHpiRequest, error) {
	collection := db.GetCollection("hpi")
	var updated dtoHpi.ActiveHpiRequest

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

// Biomarker

func CreateBiomarker(ctx context.Context, id primitive.ObjectID, biomarker dtoHpi.CreateBiomarkerRequest) (dtoHpi.CreateBiomarkerRequest, error) {
	collection := db.GetCollection("hpi")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$push": bson.M{"biomarker": biomarker},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return biomarker, err
}

func GetBiomarkerByName(ctx context.Context, name string) (*dtoHpi.CreateBiomarkerRequest, error) {
	collection := db.GetCollection("hpi")
	var biomarker dtoHpi.CreateBiomarkerRequest
	err := collection.FindOne(ctx, bson.M{"biomarker.name": name}).Decode(&biomarker)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &biomarker, nil
}

func GetBiomarkerById(ctx context.Context, id primitive.ObjectID) (*dtoHpi.CreateBiomarkerRequest, error) {
	collection := db.GetCollection("hpi")
	var hpi struct {
		Biomarker []dtoHpi.CreateBiomarkerRequest `bson:"biomarker"`
	}
	err := collection.FindOne(ctx, bson.M{"biomarker._id": id}).Decode(&hpi)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	for _, biomarker := range hpi.Biomarker {
		if biomarker.Id == id {
			return &biomarker, nil
		}
	}

	return nil, nil
}

func GetActiveBiomarkerById(ctx context.Context, id primitive.ObjectID) (*dtoHpi.CreateBiomarkerRequest, error) {
	collection := db.GetCollection("hpi")
	var biomarker dtoHpi.CreateBiomarkerRequest

	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&biomarker)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &biomarker, nil
}

func GetAllBiomarkers(ctx context.Context, id primitive.ObjectID) ([]dtoHpi.CreateBiomarkerRequest, error) {
	collection := db.GetCollection("hpi")
	var result []dtoHpi.CreateBiomarkerRequest

	filter := bson.M{
		"_id":        id,
		"deleted_at": bson.M{"$eq": nil},
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var hpi struct {
			Biomarker []dtoHpi.CreateBiomarkerRequest `bson:"biomarker"`
		}

		if err := cursor.Decode(&hpi); err != nil {
			return nil, err
		}
		for _, biomarker := range hpi.Biomarker {
			if biomarker.DeletedAt == nil {
				result = append(result, biomarker)
			}
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetAllNonActiveBiomarkers(ctx context.Context, idHpi primitive.ObjectID) ([]dtoHpi.CreateBiomarkerRequest, error) {
	collection := db.GetCollection("hpi")
	var result []dtoHpi.CreateBiomarkerRequest

	filter := bson.M{
		"_id": idHpi,
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var hpi struct {
			Biomarker []dtoHpi.CreateBiomarkerRequest `bson:"biomarker"`
		}

		if err := cursor.Decode(&hpi); err != nil {
			return nil, err
		}
		for _, biomarker := range hpi.Biomarker {
			if biomarker.DeletedAt != nil {
				result = append(result, biomarker)
			}
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateBiomarkerById(ctx context.Context, id primitive.ObjectID, updateData dtoHpi.UpdateBiomarkerRequest) (*dtoHpi.UpdateBiomarkerRequest, error) {
	collection := db.GetCollection("hpi")
	var updated dtoHpi.UpdateBiomarkerRequest

	filter := bson.M{"biomarker._id": id}

	update := bson.M{
		"$set": bson.M{
			"biomarker.$.name":       updateData.Name,
			"biomarker.$.updated_by": updateData.UpdatedBy,
			"biomarker.$.updated_at": updateData.UpdatedAt,
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

func ActiveBiomarkerById(ctx context.Context, id primitive.ObjectID, updateData dtoHpi.ActiveBiomarkerRequest) (*dtoHpi.ActiveBiomarkerRequest, error) {
	collection := db.GetCollection("hpi")
	var updated dtoHpi.ActiveBiomarkerRequest

	filter := bson.M{"biomarker._id": id}

	update := bson.M{
		"$set": bson.M{
			"biomarker.$.deleted_at": updateData.DeletedAt,
			"biomarker.$.updated_by": updateData.UpdatedBy,
			"biomarker.$.updated_at": updateData.UpdatedAt,
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

// Under

func CreateUnder(ctx context.Context, biomarkerId primitive.ObjectID, under dtoHpi.CreateUnderRequest) (dtoHpi.CreateUnderRequest, error) {
	collection := db.GetCollection("hpi")
	filter := bson.M{"biomarker._id": biomarkerId}
	update := bson.M{
		"$set": bson.M{"biomarker.$.under": under},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return under, err
}

func GetUnderByBiomarkerId(ctx context.Context, biomarkerId primitive.ObjectID) (*dtoHpi.CreateUnderRequest, error) {
	collection := db.GetCollection("hpi")
	var hpi struct {
		Biomarker []struct {
			Under dtoHpi.CreateUnderRequest `bson:"under"`
		} `bson:"biomarker"`
	}
	err := collection.FindOne(ctx, bson.M{"biomarker._id": biomarkerId}).Decode(&hpi)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	for _, biomarker := range hpi.Biomarker {
		if biomarker.Under != (dtoHpi.CreateUnderRequest{}) {
			return &biomarker.Under, nil
		}
	}

	return nil, nil
}

func UpdateUnderByBiomarkerId(ctx context.Context, biomarkerId primitive.ObjectID, updateData dtoHpi.UpdateUnderRequest) (*dtoHpi.UpdateUnderRequest, error) {
	collection := db.GetCollection("hpi")
	var updated dtoHpi.UpdateUnderRequest

	filter := bson.M{"biomarker._id": biomarkerId}

	update := bson.M{
		"$set": bson.M{},
	}

	if updateData.Value != "" {
		update["$set"].(bson.M)["biomarker.$.under.value"] = updateData.Value
	}
	if updateData.Unit != "" {
		update["$set"].(bson.M)["biomarker.$.under.unit"] = updateData.Unit
	}
	if updateData.Excercise != "" {
		update["$set"].(bson.M)["biomarker.$.under.excercise"] = updateData.Excercise
	}
	if updateData.Nutrision != "" {
		update["$set"].(bson.M)["biomarker.$.under.nutrision"] = updateData.Nutrision
	}
	update["$set"].(bson.M)["biomarker.$.under.updated_by"] = updateData.UpdatedBy
	update["$set"].(bson.M)["biomarker.$.under.updated_at"] = updateData.UpdatedAt

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
func CreateOver(ctx context.Context, biomarkerId primitive.ObjectID, over dtoHpi.CreateOverRequest) (dtoHpi.CreateOverRequest, error) {
	collection := db.GetCollection("hpi")
	filter := bson.M{"biomarker._id": biomarkerId}
	update := bson.M{
		"$set": bson.M{"biomarker.$.over": over},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return over, err
}

func GetOverByBiomarkerId(ctx context.Context, biomarkerId primitive.ObjectID) (*dtoHpi.CreateOverRequest, error) {
	collection := db.GetCollection("hpi")
	var hpi struct {
		Biomarker []struct {
			Over dtoHpi.CreateOverRequest `bson:"over"`
		} `bson:"biomarker"`
	}
	err := collection.FindOne(ctx, bson.M{"biomarker._id": biomarkerId}).Decode(&hpi)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	for _, biomarker := range hpi.Biomarker {
		if biomarker.Over != (dtoHpi.CreateOverRequest{}) {
			return &biomarker.Over, nil
		}
	}

	return nil, nil
}

func UpdateOverByBiomarkerId(ctx context.Context, biomarkerId primitive.ObjectID, updateData dtoHpi.UpdateOverRequest) (*dtoHpi.UpdateOverRequest, error) {
	collection := db.GetCollection("hpi")
	var updated dtoHpi.UpdateOverRequest

	filter := bson.M{"biomarker._id": biomarkerId}

	update := bson.M{
		"$set": bson.M{},
	}

	if updateData.Value != "" {
		update["$set"].(bson.M)["biomarker.$.over.value"] = updateData.Value
	}
	if updateData.Unit != "" {
		update["$set"].(bson.M)["biomarker.$.over.unit"] = updateData.Unit
	}
	if updateData.Excercise != "" {
		update["$set"].(bson.M)["biomarker.$.over.excercise"] = updateData.Excercise
	}
	if updateData.Nutrision != "" {
		update["$set"].(bson.M)["biomarker.$.over.nutrision"] = updateData.Nutrision
	}
	update["$set"].(bson.M)["biomarker.$.over.updated_by"] = updateData.UpdatedBy
	update["$set"].(bson.M)["biomarker.$.over.updated_at"] = updateData.UpdatedAt

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

func UpdateLifeStyleById(ctx context.Context, id primitive.ObjectID, updateData dtoHpi.UpdateLifestyleRequest) (*dtoHpi.UpdateLifestyleRequest, error) {
	collection := db.GetCollection("hpi")
	var updated dtoHpi.UpdateLifestyleRequest

	filter := bson.M{"biomarker._id": id}

	update := bson.M{
		"$set": bson.M{
			"biomarker.$.lifestyle":  updateData.Lifestyle,
			"biomarker.$.updated_by": updateData.UpdatedBy,
			"biomarker.$.updated_at": updateData.UpdatedAt,
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

// func HpiResults(ctx context.Context, id primitive.ObjectID) (dtoHpi.HpiResult ,error){

// }

func CreateHpiResult(ctx context.Context, HpiPayload dtoHpi.HpiResult) (dtoHpi.HpiResult, error) {
	collection := db.GetCollection("hpi_result")
	_, err := collection.InsertOne(ctx, HpiPayload)
	return HpiPayload, err
}

func GetHpiResultByIdUser(ctx context.Context, idUser primitive.ObjectID) (*[]dtoHpi.HpiResult, error) {
	collection := db.GetCollection("hpi_result")
	var hpiResults []dtoHpi.HpiResult

	cursor, err := collection.Find(ctx, bson.M{"_id_user": idUser})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var hpiResult dtoHpi.HpiResult
		if err := cursor.Decode(&hpiResult); err != nil {
			return nil, err
		}
		hpiResults = append(hpiResults, hpiResult)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &hpiResults, nil
}
