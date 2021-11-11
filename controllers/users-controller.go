package controllers

import (
	"github.com/avnigenc/go-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MeController(c *gin.Context)  {
	userId, _ := c.Get("user_id")

	user, err := services.GetUserById(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User": user,
	})
	return
}

func UpdateUserController(c *gin.Context)  {

}