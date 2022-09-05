package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/avnigenc/go-api/controllers"
	models "github.com/avnigenc/go-api/models/api"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHealthController(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/health", controllers.HealthController)

	req, _ := http.NewRequest("GET", "/api/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestIndexController(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api", controllers.IndexController)

	req, _ := http.NewRequest("GET", "/api", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var genericResponse models.GenericResponse
	json.Unmarshal(w.Body.Bytes(), &genericResponse)

	assert.Equal(t, genericResponse.Result, "pong")
	assert.Equal(t, http.StatusOK, w.Code)
}
