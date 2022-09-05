package services

import (
	"errors"
	"fmt"

	models "github.com/avnigenc/go-api/models/api"
	repository "github.com/avnigenc/go-api/repository"
	"github.com/avnigenc/go-api/utils"
)

func Authenticate(loginModel models.LoginRequestModel) (string, error) {
	user, err := repository.GetByEmail(loginModel.Email)

	if err != nil {
		fmt.Println("[UserService]: user not found: ", loginModel.Email)
		return "", errors.New("user not found")
	}

	if !utils.CheckPasswordHash(loginModel.Password, user.Password) {
		fmt.Println("[UserService]: password wrong: ", loginModel.Email)
		return "", errors.New("password wrong")
	}

	token, _ := utils.GenerateAccessToken(user.Id)
	return token, nil
}
