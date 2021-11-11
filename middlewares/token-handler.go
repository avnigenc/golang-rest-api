package middlewares

import (
	"github.com/avnigenc/go-api/shared"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
	"strings"
)

func TokenHandler(c *gin.Context) {

	var cfg shared.Config
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal("config error")
	}

	bearerToken := c.Request.Header["Authorization"][0]
	if bearerToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "token not found",
		})
		c.Abort()
		return
	}
	tokenString := strings.Split(bearerToken, " ")[1]

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JwtSecret), nil
	})

	if err != nil {
		if err.Error() == "Token is expired" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token expired",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "token invalid",
		})
		c.Abort()
		return
	}

	for key, val := range claims {
		if key == "user_id" {
			c.Set("user_id", val)
			return
		}
	}
}
