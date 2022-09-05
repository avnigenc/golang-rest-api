package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/google/uuid"
)

func GenerateAccessToken(userID primitive.ObjectID) (string, error) {
	jwtTimeString := os.Getenv("JwtExpireTime")
	tokenIssuerString := os.Getenv("TokenIssuer")
	tokenAudienceString := os.Getenv("TokenAudience")
	jwtSecretString := os.Getenv("JwtSecret")

	jwtTimeDuration, _ := time.ParseDuration(jwtTimeString)

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(jwtTimeDuration).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = userID
	claims["iss"] = tokenIssuerString
	claims["aud"] = tokenAudienceString
	claims["jti"] = uuid.New().String()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(jwtSecretString))
	if err != nil {
		return "", err
	}

	return token, nil
}
