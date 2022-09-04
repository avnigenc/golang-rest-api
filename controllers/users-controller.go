package controllers

import (
	"net/http"

	error "github.com/avnigenc/go-api/error"
	models "github.com/avnigenc/go-api/models/api"
	business "github.com/avnigenc/go-api/models/business"
	"github.com/avnigenc/go-api/services"
	"github.com/gin-gonic/gin"
)

func MeController(c *gin.Context) {
	userId, _ := c.Get("user_id")

	user, err := services.GetUserById(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.GenericResponse{
			Result: nil,
			Error:  error.NewBadRequestError("user not found!"),
		})
		return
	}

	c.JSON(http.StatusOK, &models.GenericResponse{
		Result: business.User{
			Id:    user.Id,
			Email: user.Email,
		},
		Error: nil,
	})
	return
}

func UpdateUserController(c *gin.Context) {

}
