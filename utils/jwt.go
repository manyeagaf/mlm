package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// SecretKey for signing JWTs
var SecretKey = []byte("your_secret_key") // Change this to a secure key

// GenerateJWT generates a token for a user
func GenerateJWT(userID uint, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // Token expiration time (24 hours)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

// ValidateJWT validates and parses a token
func ValidateJWT(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, err
}
