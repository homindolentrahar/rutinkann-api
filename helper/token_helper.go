package helper

import (
	"errors"
	"os"
	"strconv"
	"time"

	"com.homindolentrahar.rutinkann-api/model"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userClaim *model.UserClaim) (string, error) {
	// Set expired token to a week
	expiredTime := time.Now().Add(7 * (24 * time.Hour))
	secretKey := os.Getenv("APP_SECRET_KEY")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        strconv.Itoa(userClaim.ID),
		Issuer:    os.Getenv("APP_NAME"),
		Subject:   userClaim.Username,
		ExpiresAt: jwt.NewNumericDate(expiredTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	return claims.SignedString([]byte(secretKey))
}

func VerifyToken(secretKey string, tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid authentication token")
	}

	return token, nil
}
