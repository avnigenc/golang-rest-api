package controllers

import (
	"net/http"

	error "github.com/avnigenc/go-api/error"
	models "github.com/avnigenc/go-api/models/api"

	"github.com/avnigenc/go-api/services"
	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var loginRequestModel models.LoginRequestModel

	if err := c.ShouldBindJSON(&loginRequestModel); err != nil {
		c.JSON(http.StatusBadRequest, &models.GenericResponse{
			Result: nil,
			Error:  error.NewBadRequestError(err.Error()),
		})
		return
	}

	token, err := services.Authenticate(loginRequestModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.GenericResponse{
			Result: nil,
			Error:  error.NewBadRequestError(err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, &models.GenericResponse{
		Result: models.LoginResponseModel{
			Token: token,
		},
		Error: nil,
	})
	return
}

func RegisterController(c *gin.Context) {
	var registerModel models.RegisterRequestModel

	if err := c.ShouldBindJSON(&registerModel); err != nil {
		c.JSON(http.StatusBadRequest, &models.GenericResponse{
			Result: nil,
			Error:  error.NewBadRequestError(err.Error()),
		})
		return
	}

	user, err := services.CreateUser(registerModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.GenericResponse{
			Result: nil,
			Error:  error.NewBadRequestError(err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, &models.GenericResponse{
		Result: models.RegisterResponseModel{
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
		Error: nil,
	})
	return
}
