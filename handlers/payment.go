package handlers

import (
	"net/http"
	"server/db"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePaymentHandler(c *gin.Context) {
	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if payment.ID.IsZero() {
		payment.ID = primitive.NewObjectID()
	}

	payment.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	createdBy := c.GetHeader("Authorization")
	if createdBy == "" {
		createdBy = "system"
	}
	payment.CreatedBy = createdBy
	collection := db.GetCollection("payments")
	_, err := collection.InsertOne(c, payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Payment created successfully", "data": payment})
}

func GetPaymentHandler(c *gin.Context) {
	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	collection := db.GetCollection("payments")
	filter := bson.M{"_id": objID, "deleted_at": bson.M{"$exists": false}}
	var payment models.Payment

	err = collection.FindOne(c, filter).Decode(&payment)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment})
}

func UpdatePaymentHandler(c *gin.Context) {
	id := c.Param("id")
	var input models.Payment

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBy := c.GetHeader("Authorization")
	if updatedBy == "" {
		updatedBy = "system" // Jika tidak ada token, gunakan default
	}

	collection := db.GetCollection("payments")
	filter := bson.M{"_id": objID, "deleted_at": bson.M{"$exists": false}}
	update := bson.M{
		"$set": bson.M{
			"_id_user":    input.UserID,
			"_id_product": input.ProductID,
			"updated_at":  primitive.NewDateTimeFromTime(time.Now()),
			"updated_by":  "system",
		},
	}

	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment updated successfully"})
}

func DeletePaymentHandler(c *gin.Context) {
	id := c.Param("id")

	// Konversi ID string ke ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Cari data dan update field `deletedAt`
	collection := db.GetCollection("payments")
	filter := bson.M{"_id": objID, "deleted_at": bson.M{"$exists": false}}
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.NewDateTimeFromTime(time.Now()),
		},
	}

	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment soft deleted successfully"})
}
