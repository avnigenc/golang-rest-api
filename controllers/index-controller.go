package controllers

import (
	"net/http"
	"time"

	api_models "github.com/avnigenc/go-api/models/api"
	"github.com/gin-gonic/gin"
)

func IndexController(c *gin.Context) {
	c.JSON(http.StatusOK, &api_models.GenericResponse{
		Result: "pong",
		Error:  nil,
	})

	return
}

func HealthController(c *gin.Context) {
	c.JSON(http.StatusOK, &api_models.GenericResponse{
		Result: api_models.HealthResponse{
			Time:    int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond),
			Message: "OK",
		},
		Error: nil,
	})
	return
}
