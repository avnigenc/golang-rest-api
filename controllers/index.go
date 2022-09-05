package controllers

import (
	"net/http"
	"time"

	models "github.com/avnigenc/go-api/models/api"
	"github.com/gin-gonic/gin"
)

func IndexController(c *gin.Context) {
	c.JSON(http.StatusOK, &models.GenericResponse{
		Result: "pong",
		Error:  nil,
	})
}

func HealthController(c *gin.Context) {
	c.JSON(http.StatusOK, &models.GenericResponse{
		Result: models.HealthResponse{
			Time:    int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond),
			Message: "OK",
		},
		Error: nil,
	})
}
