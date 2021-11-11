package utils

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"
)

func GenerateAccessToken(userId primitive.ObjectID) (string, error) {
	var err error

	jwtTimeString := os.Getenv("JwtExpireTime")
	jwtTimeDuration, _ := time.ParseDuration(jwtTimeString)

	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(jwtTimeDuration).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(os.Getenv("JwtSecret")))
	if err != nil {
		return "", err
	}

	return token, nil
}

