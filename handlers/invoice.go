package handlers

import (
	"context"
	"net/http"
	"server/dto"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateInvoiceHandler: Handler untuk membuat invoice baru
func CreateInvoiceHandler(c *gin.Context) {
	var invoiceDTO dto.InvoiceDTO
	if err := c.ShouldBindJSON(&invoiceDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoiceDTO.CreatedAt = time.Now()
	invoiceDTO.UpdatedAt = time.Now()

	invoiceID, err := repositories.CreateInvoice(context.Background(), invoiceDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": invoiceID})
}

// GetInvoiceByIDHandler: Handler untuk mendapatkan invoice berdasarkan ID
func GetInvoiceByIDHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak boleh kosong"})
		return
	}

	invoice, err := repositories.GetInvoiceByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if invoice.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invoice tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

// UpdateInvoiceHandler: Handler untuk memperbarui invoice berdasarkan ID
func UpdateInvoiceHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak boleh kosong"})
		return
	}

	var invoiceDTO dto.InvoiceDTO
	if err := c.ShouldBindJSON(&invoiceDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoiceDTO.UpdatedAt = time.Now()

	err := repositories.UpdateInvoice(context.Background(), id, invoiceDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// DeleteInvoiceHandler: Handler untuk menghapus invoice berdasarkan ID
func SoftDeleteInvoiceHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	err := repositories.SoftDeleteInvoice(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
