package middlewares

import (
	"net/http"
	"os"
	"strings"

	apiError "github.com/avnigenc/go-api/error"
	models "github.com/avnigenc/go-api/models/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TokenHandler(c *gin.Context) {
	jwtSecretString := os.Getenv("JwtSecret")

	if len(c.Request.Header["Authorization"]) == 0 {
		c.JSON(http.StatusBadRequest, &models.GenericResponse{
			Result: nil,
			Error:  apiError.NewBadRequestError("token not found!"),
		})
		c.Abort()
		return
	}

	bearerToken := c.Request.Header["Authorization"][0]
	if bearerToken == "" {
		c.JSON(http.StatusBadRequest, &models.GenericResponse{
			Result: nil,
			Error:  apiError.NewBadRequestError("token malformed"),
		})
		c.Abort()
		return
	}

	tokenString := strings.Split(bearerToken, " ")[1]

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretString), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, &models.GenericResponse{
			Result: nil,
			Error:  apiError.NewBadRequestError(err.Error()),
		})
		c.Abort()
		return
	}

	for key, val := range claims {
		if key == "sub" {
			c.Set("user_id", val)
			return
		}
	}
}
