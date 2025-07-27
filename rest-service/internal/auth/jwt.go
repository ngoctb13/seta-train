package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your-secret-key")

func GenerateJWT(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func GetUserIDFromToken(tokenString string) (string, error) {
	claims, err := ParseJWT(tokenString)
	if err != nil {
		return "", err
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("user_id not found in token")
	}
	return userID, nil
}

func ParseAndValidateJWT(tokenString string) (string, string, error) {
	if tokenString == "" {
		return "", "", errors.New("empty token")
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	claims, err := ParseJWT(tokenString)
	if err != nil {
		return "", "", err
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", "", errors.New("user_id not found in token")
	}
	role, _ := claims["role"].(string)
	return userID, role, nil
}
