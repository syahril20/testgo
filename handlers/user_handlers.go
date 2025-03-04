package handlers

import (
	"context"
	"log"
	"net/http"
	"server/db"
	"server/dto"
	"server/models"
	"server/repositories"
	"time"

	dtoResult "server/dto/result"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword melakukan hashing password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CreateUserHandler membuat user baru berdasarkan request
func CreateUserHandler(c *gin.Context) {
	var req dto.CreateUserRequest

	// Bind JSON request ke struct CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi format birth date
	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid birth date format. Use YYYY-MM-DD"})
		return
	}

	// Hash password
	hashedPassword, err := HashPassword(req.Password.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Menentukan waktu saat ini
	currentTime := time.Now()

	// Membuat objek user baru
	user := models.User{
		Email: models.UserEmail{
			Value: req.Email.Value,
		},
		Level: req.Level,
		Password: models.UserPassword{
			Value: hashedPassword,
		},
		Data: models.UserData{
			Name:      req.Name,
			BirthDate: birthDate,
			Gender:    req.Gender,
		},
		Ring: models.UserRing{
			Size:       req.Ring.Size,
			Color:      req.Ring.Color,
			Connection: req.Ring.Connection,
		},
		Personal: models.UserPersonal{
			Health: models.UserHealth{
				Allergies:     req.Personal.Health.Allergies,
				Diseases:      req.Personal.Health.Diseases,
				Goals:         req.Personal.Health.Goals,
				BloodType:     req.Personal.Health.BloodType,
				Issues:        []string{},
				SpecificGoals: []string{},
			},
			Physical: models.UserPhysical{
				Height:    req.Personal.Physical.Height,
				Weight:    req.Personal.Physical.Weight,
				Abdominal: req.Personal.Physical.Abdominal,
				Waist:     req.Personal.Physical.Waist,
				Unit:      req.Personal.Physical.Unit,
			},
			Habit: models.UserHabit{
				Smoke:   req.Personal.Habit.Smoke,
				Alcohol: req.Personal.Habit.Alcohol,
			},
		},
		CreatedAt: currentTime,
		CreatedBy: "SYSTEM",
		UpdatedAt: currentTime,
		UpdatedBy: "SYSTEM",
	}

	// Membuka konteks untuk operasi MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Memanggil fungsi untuk menyimpan user baru ke database
	if err := repositories.CreateUser(ctx, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menyampaikan respon sukses setelah berhasil membuat user
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// GetUserByIDHandler mengambil data user berdasarkan ID
func GetUserByIDHandler(c *gin.Context) {
	id := c.Param("id")
	activity, err := repositories.GetUserByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func GetUserHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
		return
	}

	claims, _ := ValidateToken(c, authHeader)
	user, err := repositories.GetUserByEmailV2(context.Background(), claims.UserData.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, dtoResult.ErrorResult{
			Code:    http.StatusNotFound,
			Message: "user not found"})
		return
	}

	if claims != nil {
		c.JSON(http.StatusOK, dtoResult.SuccessResult{
			Code:    http.StatusOK,
			Message: "success",
			Data:    user})
		return
	}
}

// UpdateUserHandler memperbarui data user berdasarkan ID
func UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")
	var updateData dto.UpdateUserRequest

	// Bind data dari request JSON
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateFields := bson.M{}

	log.Println("Received update request:", updateData)

	if updateData.Data.BirthDate != "" {
		parsedDate, err := time.Parse("2006-01-02", updateData.Data.BirthDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, please use YYYY-MM-DD"})
			return
		}
		updateFields["data.birth_date"] = parsedDate
		log.Println("Updated birth_date:", parsedDate)
	}

	// Update acknowledged_tos jika diberikan
	if updateData.Data.AcknowledgedTOS {
		updateFields["data.acknowledged_tos"] = true
		log.Println("Updated acknowledged_tos:", true)
	}

	// Cek dan update first_login jika nilai yang dikirimkan adalah true atau false
	updateFields["data.first_login"] = updateData.Data.FirstLogin
	log.Println("Updated first_login:", updateData.Data.FirstLogin)

	// Update password jika diberikan
	if updateData.Password.Value != "" {
		hashedPassword, err := HashPassword(updateData.Password.Value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		updateFields["password.value"] = hashedPassword
		log.Println("Updated password:", hashedPassword)
	}

	// Set waktu update dan siapa yang meng-update
	updateFields["updated_at"] = time.Now()
	updateFields["updated_by"] = "ADMIN"

	// Konversi ID menjadi ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Membuka konteks untuk operasi MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Melakukan update di database
	collection := db.GetCollection("users")
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updateFields}

	// Lakukan update ke database
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Ambil data yang sudah diperbarui
	var updatedUser dto.UpdateUserRequest
	err = collection.FindOne(ctx, filter).Decode(&updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated user"})
		return
	}

	// Debugging: Print hasil update
	log.Println("Updated user data:", updatedUser)

	// Kirimkan respons dengan data yang diperbarui
	c.JSON(http.StatusOK, gin.H{
		"message":      "User updated successfully",
		"updated_user": updatedUser,
	})
}
