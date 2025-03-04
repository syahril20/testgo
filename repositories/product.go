package repositories

import (
	"context"
	"server/db"
	dtoProduct "server/dto/product"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateProduct(ctx context.Context, product dtoProduct.CreateProductRequest) (dtoProduct.CreateProductRequest, error) {
	collection := db.GetCollection("product")
	_, err := collection.InsertOne(ctx, product)
	return product, err
}

func GetAllProduct(ctx context.Context) ([]dtoProduct.CreateProductRequest, error) {
	collection := db.GetCollection("product")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Inisialisasi slice produk sebagai array kosong
	product := []dtoProduct.CreateProductRequest{}

	for cursor.Next(ctx) {
		var products dtoProduct.CreateProductRequest
		if err := cursor.Decode(&products); err != nil {
			return nil, err
		}
		product = append(product, products)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return product, nil
}

func GetProductByID(ctx context.Context, id primitive.ObjectID) (*dtoProduct.ProductResponse, error) {
	collection := db.GetCollection("product")
	var product dtoProduct.ProductResponse

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &product, err
}

func GetProductByNameID(ctx context.Context, id primitive.ObjectID, name string) (*dtoProduct.ProductResponse, error) {
	collection := db.GetCollection("product")
	var product dtoProduct.ProductResponse

	filter := bson.M{
		"$and": []bson.M{
			{"_id": id},    // ID produk utama
			{"name": name}, // Nama sub produk
		},
	}

	err := collection.FindOne(ctx, filter).Decode(&product)
	return &product, err
}

func GetProductByName(ctx context.Context, name string) (*dtoProduct.ProductResponse, error) {
	collection := db.GetCollection("product")
	var product dtoProduct.ProductResponse
	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

// Sub Product
func CreateSubProduct(ctx context.Context, objId primitive.ObjectID, subProduct dtoProduct.CreateSubProductRequest) (dtoProduct.CreateSubProductRequest, error) {
	collection := db.GetCollection("product")
	filter := bson.M{"_id": objId}
	update := bson.M{
		"$push": bson.M{"sub_product": subProduct},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return subProduct, err
}

func GetSubProductByName(ctx context.Context, name string) (*dtoProduct.CreateSubProductRequest, error) {
	collection := db.GetCollection("product")
	var subProduct dtoProduct.CreateSubProductRequest

	filter := bson.M{"sub_product.name": name}

	err := collection.FindOne(ctx, filter).Decode(&subProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &subProduct, nil
}

func GetSubProductByNameId(ctx context.Context, idProduct primitive.ObjectID, name string) (*dtoProduct.CreateSubProductRequest, error) {
	collection := db.GetCollection("product")
	var product dtoProduct.CreateProductRequest

	// Filter pencarian
	filter := bson.M{
		"_id": idProduct,
		"sub_product": bson.M{
			"$elemMatch": bson.M{
				"name":       name,
				"deleted_at": bson.M{"$eq": nil},
			},
		},
	}

	// Cari dokumen
	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	// Filter sub_product secara manual
	for _, sp := range product.SubProduct {
		if sp.IdProduct == idProduct && sp.Name == name {
			return &sp, nil
		}
	}

	return nil, nil
}

func GetSubProductById(ctx context.Context, subProductID primitive.ObjectID) (*dtoProduct.CreateSubProductRequest, error) {
	collection := db.GetCollection("product")
	var subProduct dtoProduct.CreateSubProductRequest

	filter := bson.M{"sub_product._id": subProductID}

	err := collection.FindOne(ctx, filter).Decode(&subProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &subProduct, nil
}

// Addons
func CreateAddons(ctx context.Context, subProductID primitive.ObjectID, addons dtoProduct.CreateAddonsRequest) error {
	collection := db.GetCollection("product")

	filter := bson.M{
		"sub_product._id": subProductID,
	}

	update := bson.M{
		"$push": bson.M{"sub_product.$.addons": addons},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func GetAddonsByName(ctx context.Context, name string) (*dtoProduct.CreateAddonsRequest, error) {
	collection := db.GetCollection("product")
	var subProduct dtoProduct.CreateAddonsRequest

	// Filter pencarian
	filter := bson.M{
		"sub_product.addons": bson.M{
			"$elemMatch": bson.M{
				"name":       name,
				"deleted_at": bson.M{"$eq": nil},
			},
		},
	}

	err := collection.FindOne(ctx, filter).Decode(&subProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &subProduct, nil
}
