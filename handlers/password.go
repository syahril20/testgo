package handlers

import (
	"context"
	"net/http"
	"server/dto"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UpgradePasswordHandler(c *gin.Context) {
	// Ambil email dari parameter URL
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email parameter is required"})
		return
	}

	var req dto.UpgradePasswordRequest
	// Bind JSON body ke struct DTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	// Buat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ambil data user berdasarkan email
	user, err := repositories.GetUserByID(ctx, email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Verifikasi password lama
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password.Value), []byte(req.OldPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Old password is incorrect"})
		return
	}

	// Hash password baru
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash new password"})
		return
	}

	// Gunakan user.ID langsung tanpa Hex()
	err = repositories.UpdatePassword(ctx, user.ID, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password", "details": err.Error()})
		return
	}

	// Berhasil mengubah password
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
