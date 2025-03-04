package pkg

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"os"
	"time"
)

// LoadEnv untuk memuat variabel lingkungan dari file .env
func LoadEnv() error {
	return godotenv.Load() // Pastikan Anda telah memuat file .env
}

// SecretKey diambil dari variabel lingkungan
var SecretKey = os.Getenv("SECRET_KEY") // Mengambil SecretKey dari .env

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email, userID string) (string, error) {
	claims := Claims{
		ID:    userID,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SecretKey))
}
