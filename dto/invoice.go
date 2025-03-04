package dto

import "time"

type InvoiceDTO struct {
	ID        string    `json:"id"`
	PaymentID string    `json:"payment_id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
