package controllers

import (
	"github.com/avnigenc/go-api/models/api"
	"github.com/avnigenc/go-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginController(c *gin.Context) {
	var loginRequestModel api_models.LoginRequestModel

	if err := c.ShouldBindJSON(&loginRequestModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	token, err := services.Authenticate(loginRequestModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"email": loginRequestModel.Email,
			"token": token,
		},
		"statusCode": http.StatusOK,
		"message": "user logged in successfully",
	})
	return
}

func RegisterController(c *gin.Context) {
	var registerModel api_models.RegisterRequestModel

	if err := c.ShouldBindJSON(&registerModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	user, err := services.CreateUser(registerModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H {
			"Id": 	user.Id,
			"FirstName": 	user.FirstName,
			"LastName": 	user.LastName,
			"Email": 		user.Email,
		},
		"statusCode": http.StatusOK,
		"message": "user created",
	})
	return
}