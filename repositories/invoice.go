package repositories

import (
	"context"
	"server/db"
	"server/dto"
	"server/models"
	"server/pkg" // Import pkg untuk ParseObjectID
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateInvoice: Menambahkan invoice baru
func CreateInvoice(ctx context.Context, invoiceDTO dto.InvoiceDTO) (string, error) {
	collection := db.GetCollection("invoices")

	// Konversi DTO ke Model
	paymentID, err := pkg.ParseObjectID(invoiceDTO.PaymentID)
	if err != nil {
		return "", err
	}

	invoice := models.Invoice{
		ID:        primitive.NewObjectID(),
		PaymentID: paymentID,
		CreatedAt: invoiceDTO.CreatedAt,
		CreatedBy: invoiceDTO.CreatedBy,
		UpdatedAt: invoiceDTO.UpdatedAt,
		UpdatedBy: invoiceDTO.UpdatedBy,
	}

	_, err = collection.InsertOne(ctx, invoice)
	if err != nil {
		return "", err
	}

	return invoice.ID.Hex(), nil
}

func GetInvoiceByID(ctx context.Context, id string) (dto.InvoiceDTO, error) {
	collection := db.GetCollection("invoices")

	objectID, err := pkg.ParseObjectID(id)
	if err != nil {
		return dto.InvoiceDTO{}, err
	}

	var invoice models.Invoice
	err = collection.FindOne(ctx, bson.M{"_id": objectID, "deleted_at": nil}).Decode(&invoice)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dto.InvoiceDTO{}, nil // Not found
		}
		return dto.InvoiceDTO{}, err
	}

	// Convert model to DTO
	return convertInvoiceToDTO(invoice), nil
}

func UpdateInvoice(ctx context.Context, id string, invoiceDTO dto.InvoiceDTO) error {
	collection := db.GetCollection("invoices")

	objectID, err := pkg.ParseObjectID(id)
	if err != nil {
		return err
	}

	// Check if the invoice exists and is not soft-deleted
	var existingInvoice models.Invoice
	err = collection.FindOne(ctx, bson.M{"_id": objectID, "deleted_at": nil}).Decode(&existingInvoice)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return mongo.ErrNoDocuments // Invoice not found or already deleted
		}
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"payment_id": invoiceDTO.PaymentID,
			"updated_at": invoiceDTO.UpdatedAt,
			"updated_by": invoiceDTO.UpdatedBy,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID, "deleted_at": nil}, update)
	return err
}

// DeleteInvoice: Menghapus invoice berdasarkan ID
func SoftDeleteInvoice(ctx context.Context, id string) error {
	collection := db.GetCollection("invoices")

	objectID, err := pkg.ParseObjectID(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID, "deleted_at": nil}, update)
	return err
}

// Helper function: Konversi Model ke DTO
func convertInvoiceToDTO(invoice models.Invoice) dto.InvoiceDTO {
	return dto.InvoiceDTO{
		ID:        invoice.ID.Hex(),
		PaymentID: invoice.PaymentID.Hex(),
		CreatedAt: invoice.CreatedAt,
		CreatedBy: invoice.CreatedBy,
		UpdatedAt: invoice.UpdatedAt,
		UpdatedBy: invoice.UpdatedBy,
	}
}
