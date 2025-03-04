package handlers

import (
	"context"
	"os"
	dtoAuth "server/dto/auth"
	dto "server/dto/result"
	"server/repositories"

	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Claims struct {
	UserData UserData `json:"user_data"`
	Token    Token    `json:"token"`
	jwt.StandardClaims
}

type UserData struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
	Level     int    `json:"level"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

func Login(c *gin.Context) {
	var creds dtoAuth.AuthRequest
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := repositories.GetUserByEmail(context.Background(), creds.Email)
	if err != nil || !checkPasswordHash(creds.Password, user.Password.Value) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &Claims{
		UserData: UserData{
			Id:        user.ID,
			Name:      user.Data.Name,
			Email:     user.Email.Value,
			BirthDate: user.Data.BirthDate.Format("2006-01-02"),
			Gender:    user.Data.Gender,
			Level:     user.Level,
		},
		Token: Token{
			TokenType: "Bearer",
			ExpiresIn: 30 * 24 * 60 * 60,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	claims.Token.AccessToken = tokenString

	userId, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Member ID"})
		return
	}
	authLog := dtoAuth.AuthLog{
		Id:        primitive.NewObjectID(),
		IdUser:    userId,
		CreatedAt: time.Now(),
		CreatedBy: "System",
	}

	if err := repositories.InsertAuthLog(context.Background(), authLog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not log authentication"})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    claims})
}

func ValidateToken(c *gin.Context, authHeader string) (*Claims, error) {
	tokenString := authHeader[len("Bearer "):]

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, dto.ErrorResult{
				Code:    http.StatusUnauthorized,
				Message: "Invalid token signature",
			})
			return nil, err
		}
		c.JSON(http.StatusUnauthorized, dto.ErrorResult{
			Code:    http.StatusUnauthorized,
			Message: "Invalid token",
		})
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, dto.ErrorResult{
			Code:    http.StatusUnauthorized,
			Message: "Token has expired",
		})
		return nil, err
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, dto.ErrorResult{
			Code:    http.StatusUnauthorized,
			Message: "Invalid token",
		})
		return nil, err
	}

	return claims, nil
}
