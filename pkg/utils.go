package pkg

import "go.mongodb.org/mongo-driver/bson/primitive"

// ParseObjectID: Helper untuk parsing string ke ObjectID
func ParseObjectID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}
